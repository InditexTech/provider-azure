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

type ConfigurationObservation struct {

	// The ID of the PostgreSQL Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConfigurationParameters struct {

	// Specifies the name of the PostgreSQL Configuration, which needs to be a valid PostgreSQL configuration name. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters `json:"forProvider"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API. Sets a PostgreSQL Configuration value on a PostgreSQL Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationSpec   `json:"spec"`
	Status            ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
