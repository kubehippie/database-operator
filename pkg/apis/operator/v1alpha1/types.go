package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConditionStatus represents status of the Server or Database condition.
type ConditionStatus string

const (
	// ConditionTrue means a resource is in the condition.
	ConditionTrue ConditionStatus = "True"

	// ConditionFalse means a resource is not in the condition.
	ConditionFalse ConditionStatus = "False"

	// ConditionUnknown means controller can't decide if a Server or Database resource is in the condition or not.
	ConditionUnknown ConditionStatus = "Unknown"
)

// ServerConditionType represents condition of the Server or Database resource.
type ServerConditionType string

const (
	// Deployed means Server Deployment and Service for exposing Server or Database are created.
	Deployed ServerConditionType = "Deployed"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Server is a specification for a Server resource
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServerSpec   `json:"spec"`
	Status ServerStatus `json:"status"`
}

// ServerSpec is the spec for a Server resource
type ServerSpec struct {
}

// ServerStatus is the status for a Server resource
type ServerStatus struct {
	// Conditions indicates states of the EtcdStroageStatus,
	Conditions []ServerCondition
}

// ServerCondition contains details for the current condition of this Server instance.
type ServerCondition struct {
	// Type is the type of the condition.
	Type ServerConditionType

	// Status is the status of the condition (true, false, unknown).
	Status ConditionStatus

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string

	// Human-readable message indicating details about last transition.
	Message string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServerList is a list of Server resources
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Server `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Database is a specification for a Database resource
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec"`
	Status DatabaseStatus `json:"status"`
}

// DatabaseSpec is the spec for a Database resource
type DatabaseSpec struct {
}

// DatabaseStatus is the status for a Database resource
type DatabaseStatus struct {
	// Conditions indicates states of the EtcdStroageStatus,
	Conditions []DatabaseCondition
}

// DatabaseCondition contains details for the current condition of this Database instance.
type DatabaseCondition struct {
	// Type is the type of the condition.
	Type DatabaseConditionType

	// Status is the status of the condition (true, false, unknown).
	Status ConditionStatus

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string

	// Human-readable message indicating details about last transition.
	Message string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatabaseList is a list of Database resources
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Database `json:"items"`
}
