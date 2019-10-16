package controller

import (
	"fmt"
	"halkyon.io/operator/pkg/controller/framework"
	"halkyon.io/operator/pkg/util"
	authorizv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Role struct {
	*framework.DependentResourceHelper
}

func (res Role) Update(toUpdate runtime.Object) (bool, error) {
	return false, nil
}

func (res Role) NewInstanceWith(owner framework.Resource) framework.DependentResource {
	return newOwnedRole(owner)
}

func NewRole() Role {
	return newOwnedRole(nil)
}

func newOwnedRole(owner framework.Resource) Role {
	dependent := framework.NewDependentResource(&authorizv1.Role{}, owner)
	role := Role{DependentResourceHelper: dependent}
	dependent.SetDelegate(role)
	return role
}

func RoleName(owner framework.Resource) string {
	switch owner.(type) {
	case *Component:
		return "image-scc-privileged-role"
	case *Capability:
		return "scc-privileged-role"
	default:
		panic(fmt.Sprintf("unknown type '%s' for role owner", util.GetObjectName(owner)))
	}
}

func (res Role) Name() string {
	return RoleName(res.Owner())
}

func (res Role) Build() (runtime.Object, error) {
	c := res.Owner()
	ser := &authorizv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.GetNamespace(),
		},
		Rules: []authorizv1.PolicyRule{
			{
				APIGroups:     []string{"security.openshift.io"},
				Resources:     []string{"securitycontextconstraints"},
				ResourceNames: []string{"privileged"},
				Verbs:         []string{"use"},
			},
		},
	}

	if _, ok := c.(*Component); ok {
		ser.Rules = append(ser.Rules, authorizv1.PolicyRule{
			APIGroups: []string{"image.openshift.io"},
			Resources: []string{"imagestreams", "imagestreams/layers"},
			Verbs:     []string{"*"},
		})
	}

	return ser, nil
}

func (res Role) ShouldWatch() bool {
	return false
}
