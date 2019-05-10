package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LinkSpec defines the desired state of Link
// +k8s:openapi-gen=true
type LinkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Name                string `json:"name,omitempty"`
	TargetComponentName string `json:"targetComponentName,omitempty"`
	Kind                string `json:"kind,omitempty"`
	Ref                 string `json:"ref,omitempty"`
	// Array of env variables containing extra/additional info to be used to configure the runtime
	Envs                []Env `json:"envs,omitempty"`
}

// LinkStatus defines the observed state of Link
// +k8s:openapi-gen=true
type LinkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Phase     Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Link is the Schema for the links API
// +k8s:openapi-gen=true
type Link struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinkSpec   `json:"spec,omitempty"`
	Status LinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinkList contains a list of Link
type LinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Link `json:"items"`
}
