package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Component `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// A component represents
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ComponentSpec   `json:"spec"`
	Status            ComponentStatus `json:"status,omitempty"`
}

type ComponentSpec struct {
	// The name represents a human readable string describing from a business perspective what this component is related to
	// Example : payment-frontend, retail-backend
	Name string
	// The packagingMode refers to the archive file's type which has been used to package the code
	// Example : jar, war, ...
	PackagingMode string
	// The type is related to how the component is installed, as a pod, job, statefulset
	Type string
	// DeploymentMode indicates the strategy to be adopted to install the resources into a namespace
	// and next to create a pod. 2 strategies are currently supported; inner and outer loop
	// where outer loop refers to a build of the code and the packaging of the application into a container's image
	// while the inner loop will install a pod's running a supervisord daemon used to trigger actions such as : assemble, run, ...
	DeploymentMode string `json:"deployment,omitempty"`
	// Runtime is the framework used to start within the container the application
	// It corresponds to one of the following values: spring-boot, vertx, tornthail, nodejs
	Runtime string `json:"runtime,omitempty"`
	// To indicate if we want to expose the service out side of the cluster as a route
	ExposeService bool `json:"exposeService,omitempty"`
	// Debug port is the port exposed by the JVM for remote debugging
	DebugPort int `json:"debugPort,omitempty"`
	// The hotReload, when supported by the runtime, allows to expose an endpoint used to restart the application
	hotReload bool `json:"hotReload,omitempty"`
	// Cpu is the cpu to be assigned to the pod's running the application
	Cpu string `json:"cpu,omitempty"`
	// Cpu is the memory to be assigned to the pod's running the application
	Memory string `json:"memory,omitempty"`
	// Port is the HTTP/TCP port number used within the pod by the runtime
	Port int32 `json:"port,omitempty"`
	//
	SupervisordName string
	// The storage allows to specify the capacity and mode of the volume to be mounted for the pod
	Storage Storage `json:"storage,omitempty"`
	// The list of the images created according to the DeploymentMode to install the loop
	Images []Image `json:"image,omitempty"`
	// Array of env variables containing extra/additional info to be used to configure the runtime
	Envs []Env `json:"envs,omitempty"`
	// List of services consumed by the runtime and created as service instance from a Service Catalog
	// and next binded to a DeploymentConfig using a service binding
	Services []Service `json:"services,omitempty"`
	// The features represents a capability that it is required to have, to install to allow the component
	// to operate with by example a Prometheus backend system to collect metrics, an OpenTracing datastore
	// to centralize the traces/logs of the runtime, to deploy a servicemesh, ...
	Features []Feature `json:"features,omitempty"`
}

type ComponentStatus struct {
	Phase Phase `json:"phase,omitempty"`
}

type Feature struct {
	Name        string
	Description string
	Id          int
}

type Image struct {
	Name           string
	AnnotationCmds bool
	Repo           string
	Tag            string
	DockerImage    bool
}

type Env struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Service struct {
	Class      string `json:"class,omitempty"`
	Name       string `json:"name,omitempty"`
	Plan       string `json:"plan,omitempty"`
	ExternalId string `json:"externalid,omitempty"`
	SecretName string `json:"secretname,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	ParametersJSon string
}

type Parameter struct {
	Name  string  `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Storage struct {
	Name     string `json:"name,omitempty"`
	Capacity string `json:"capacity,omitempty"`
	Mode     string `json:"mode,omitempty"`
}

// IntegrationPhase --
type Phase string

const (
	// ComponentKind --
	ComponentKind string = "Component"

	// PhaseServiceCreation --
	PhaseServiceCreation Phase = "CreatingService"
	// PhaseBuilding --
	PhaseBuilding Phase = "Building"
	// PhaseDeploying --
	PhaseDeploying Phase = "Deploying"
	// PhaseRunning --
	PhaseRunning Phase = "Running"
	// PhaseError --
	PhaseError Phase = "Error"
)