package v1alpha

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:method=GetEphemeralContainers,verb=get,subresource=ephemeralcontainers,result=EphemeralContainers
// +genclient:method=UpdateEphemeralContainers,verb=update,subresource=ephemeralcontainers,input=EphemeralContainers,result=EphemeralContainers
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pod is the Schema of pod API
// +k8s:openapi-gen=true
type Pod struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the pod.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec PodSpec `json:"spec" yaml:"Spec"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodSpec is the special for the pod API
// +k8s:openapi-gen=true
type PodSpec struct {
	Status string `yaml:"status" json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodList contains a list of Pod
type PodList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// +optional
	Items []Pod `json:"items"`
}
