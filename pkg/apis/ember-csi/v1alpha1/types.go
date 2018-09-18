package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EmberCSIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []EmberCSI `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EmberCSI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              EmberCSISpec   `json:"spec"`
	Status            EmberCSIStatus `json:"status,omitempty"`
}

type EmberCSISpec struct {
	// Size is the size of the Controller StatefulSet
	Size 		int32  		`json:"size"`
	Config		EmberCSIConfig	`json:"config"`
}

type EmberCSIStatus struct {
	Phase string	`json:"phase"`
}


type EmberCSIConfig struct {
	BackendConfig 		string 		`json:"backendConfig"`
	PersistenceConfig	string 		`json:"persistenceConfig"`
	SystemFiles		string 		`json:"systemFiles"`
}
