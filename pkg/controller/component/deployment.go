package component

import (
	"github.com/snowdrop/component-operator/pkg/apis/component/v1alpha2"
	"github.com/snowdrop/component-operator/pkg/controller"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type deployment struct {
	base
	reconciler *ReconcileComponent // todo: remove
}

func (res deployment) NewInstanceWith(owner v1.Object) controller.DependentResource {
	return newOwnedDeployment(res.reconciler, owner)
}

func newDeployment(reconciler *ReconcileComponent) deployment {
	return newOwnedDeployment(reconciler, nil)
}

func newOwnedDeployment(reconciler *ReconcileComponent, owner v1.Object) deployment {
	dependent := newBaseDependent(&appsv1.Deployment{}, owner)
	d := deployment{
		base:       dependent,
		reconciler: reconciler,
	}
	dependent.SetDelegate(d)
	return d
}

func (res deployment) Build() (runtime.Object, error) {
	c := res.ownerAsComponent()
	if v1alpha2.BuildDeploymentMode == c.Spec.DeploymentMode {
		if err := res.reconciler.setInitialStatus(c, v1alpha2.ComponentBuilding); err != nil {
			return nil, err
		}
		return res.installBuild()
	} else if err := res.reconciler.setInitialStatus(c, v1alpha2.ComponentPending); err != nil {
		return nil, err
	}
	return res.installDev()
}

func (res deployment) Name() string {
	c := res.ownerAsComponent()
	if v1alpha2.BuildDeploymentMode == c.Spec.DeploymentMode {
		return buildNamer(c)
	}
	return defaultNamer(c)
}
