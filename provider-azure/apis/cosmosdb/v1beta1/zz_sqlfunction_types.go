/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SQLFunctionObservation struct {

	// The ID of the SQL User Defined Function.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLFunctionParameters struct {

	// Body of the User Defined Function.
	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body,omitempty"`

	// The id of the Cosmos DB SQL Container to create the SQL User Defined Function within. Changing this forces a new SQL User Defined Function to be created.
	// +crossplane:generate:reference:type=SQLContainer
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// Reference to a SQLContainer to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIdRef,omitempty" tf:"-"`

	// Selector for a SQLContainer to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIdSelector,omitempty" tf:"-"`
}

// SQLFunctionSpec defines the desired state of SQLFunction
type SQLFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLFunctionParameters `json:"forProvider"`
}

// SQLFunctionStatus defines the observed state of SQLFunction.
type SQLFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLFunction is the Schema for the SQLFunctions API. Manages an SQL User Defined Function.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLFunctionSpec   `json:"spec"`
	Status            SQLFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLFunctionList contains a list of SQLFunctions
type SQLFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLFunction `json:"items"`
}

// Repository type metadata.
var (
	SQLFunction_Kind             = "SQLFunction"
	SQLFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLFunction_Kind}.String()
	SQLFunction_KindAPIVersion   = SQLFunction_Kind + "." + CRDGroupVersion.String()
	SQLFunction_GroupVersionKind = CRDGroupVersion.WithKind(SQLFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLFunction{}, &SQLFunctionList{})
}
