package link

import (
	"context"
	"fmt"
	"halkyon.io/api/component/v1beta1"
	link "halkyon.io/api/link/v1beta1"
	"halkyon.io/operator/pkg/controller/framework"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func NewLinkManager() *LinkManager {
	return &LinkManager{}
}

type LinkManager struct {
	dependentTypes []framework.DependentResource
}

func (r *LinkManager) Helper() *framework.K8SHelper {
	return framework.GetHelperFor(r.PrimaryResourceType())
}

func (r *LinkManager) GetDependentResourcesTypes() []framework.DependentResource {
	if len(r.dependentTypes) == 0 {
		r.dependentTypes = []framework.DependentResource{newComponent()}
	}
	return r.dependentTypes
}

func (r *LinkManager) PrimaryResourceType() runtime.Object {
	return &link.Link{}
}

func (LinkManager) asLink(object runtime.Object) *Link {
	return object.(*Link)
}

func (r *LinkManager) NewFrom(name string, namespace string) (framework.Resource, error) {
	return NewLink().FetchAndInit(name, namespace, r)
}

func (r *LinkManager) CreateOrUpdate(object framework.Resource) error {
	l := r.asLink(object)

	found, err := r.fetchDeployment(l.Link)
	if err != nil {
		l.SetNeedsRequeue(true)
		return err
	}
	// Enrich the Deployment object using the information passed within the Link Spec (e.g Env Vars, EnvFrom, ...) if needed
	if containers, isModified := r.updateContainersWithLinkInfo(l.Link, found.Spec.Template.Spec.Containers); isModified {
		found.Spec.Template.Spec.Containers = containers

		if err = r.updateDeploymentWithLink(found, l); err != nil {
			// As it could be possible that we can't update the Deployment as it has been modified by another
			// process, then we will requeue
			l.SetNeedsRequeue(true)
			return err
		}

		// if the deployment has been updated, we need to update the component's status
		fetch, err := l.FetchUpdatedDependent(&v1beta1.Component{}, r.Helper())
		if err != nil {
			return fmt.Errorf("cannot retrieve associated component")
		}
		c := fetch.(*v1beta1.Component)
		compStatus := &c.Status
		compStatus.Phase = v1beta1.ComponentLinking
		compStatus.SetLinkingStatus(l.Name, v1beta1.Started, compStatus.PodName)
		compStatus.PodName = ""
		compStatus.Message = fmt.Sprintf("Initiating link %s", l.Name)
		err = r.Helper().Client.Status().Update(context.TODO(), c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LinkManager) updateContainersWithLinkInfo(l *link.Link, containers []v1.Container) ([]v1.Container, bool) {
	var isModified = false
	linkType := l.Spec.Type
	switch linkType {
	case link.SecretLinkType:
		secretName := l.Spec.Ref

		// Check if EnvFrom already exists
		// If this is the case, exit without error
		for i := 0; i < len(containers); i++ {
			var isEnvFromExist = false
			for _, env := range containers[i].EnvFrom {
				if env.SecretRef.Name == secretName {
					// EnvFrom already exists for the Secret Ref
					isEnvFromExist = true
				}
			}
			if !isEnvFromExist {
				// Add the Secret as EnvVar to the container
				containers[i].EnvFrom = append(containers[i].EnvFrom, r.addSecretAsEnvFromSource(secretName))
				isModified = true
			}
		}

	case link.EnvLinkType:
		// Check if Env already exists
		// If this is the case, exit without error
		for i := 0; i < len(containers); i++ {
			var isEnvExist = false
			for _, specEnv := range l.Spec.Envs {
				for _, env := range containers[i].Env {
					if specEnv.Name == env.Name && specEnv.Value == env.Value {
						// EnvFrom already exists for the Secret Ref
						isEnvExist = true
					}
				}
				if !isEnvExist {
					// Add the Secret as EnvVar to the container
					containers[i].Env = append(containers[i].Env, r.addKeyValueAsEnvVar(specEnv.Name, specEnv.Value))
					isModified = true
				}
			}
		}
	}

	return containers, isModified
}

func (r *LinkManager) update(d *appsv1.Deployment) error {
	helper := r.Helper()
	err := helper.Client.Update(context.TODO(), d)
	if err != nil {
		return err
	}

	helper.ReqLogger.Info("Deployment updated.")
	return nil
}

//fetchDeployment returns the deployment resource created for this instance
func (r *LinkManager) fetchDeployment(link *link.Link) (*appsv1.Deployment, error) {
	helper := r.Helper()
	deployment := &appsv1.Deployment{}
	name := link.Spec.ComponentName
	if err := helper.Client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: link.Namespace}, deployment); err == nil {
		return deployment, nil
	} else if err := helper.Client.Get(context.TODO(), types.NamespacedName{Name: name + "-build", Namespace: link.Namespace}, deployment); err == nil {
		return deployment, nil
	} else {
		helper.ReqLogger.Info("Deployment doesn't exist", "Name", name)
		return deployment, err
	}
}

func (r *LinkManager) updateDeploymentWithLink(d *appsv1.Deployment, link *Link) error {
	// Update the Deployment of the component
	if err := r.update(d); err != nil {
		r.Helper().ReqLogger.Info("Failed to update deployment.")
		return err
	}

	name := link.Spec.ComponentName
	link.SetSuccessStatus([]framework.DependentResourceStatus{}, fmt.Sprintf("linked to '%s' component", name))
	return nil
}

func (r *LinkManager) Delete(object framework.Resource) error {
	return nil
}
