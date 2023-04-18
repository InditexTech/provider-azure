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

type BasicAuthenticationObservation struct {

	// The username which can be used to authenticate to the OData endpoint.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicAuthenticationParameters struct {

	// The password associated with the username, which can be used to authenticate to the OData endpoint.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username which can be used to authenticate to the OData endpoint.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type LinkedServiceODataObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service OData.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service OData.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A basic_authentication block as defined below.
	BasicAuthentication []BasicAuthenticationObservation `json:"basicAuthentication,omitempty" tf:"basic_authentication,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service OData.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory OData Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service OData.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service OData.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The URL of the OData service endpoint.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkedServiceODataParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service OData.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service OData.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A basic_authentication block as defined below.
	// +kubebuilder:validation:Optional
	BasicAuthentication []BasicAuthenticationParameters `json:"basicAuthentication,omitempty" tf:"basic_authentication,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service OData.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service OData.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service OData.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The URL of the OData service endpoint.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// LinkedServiceODataSpec defines the desired state of LinkedServiceOData
type LinkedServiceODataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceODataParameters `json:"forProvider"`
}

// LinkedServiceODataStatus defines the observed state of LinkedServiceOData.
type LinkedServiceODataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceODataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceOData is the Schema for the LinkedServiceODatas API. Manages a Linked Service (connection) between a Database and Azure Data Factory through OData protocol.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceOData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.url)",message="url is a required parameter"
	Spec   LinkedServiceODataSpec   `json:"spec"`
	Status LinkedServiceODataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceODataList contains a list of LinkedServiceODatas
type LinkedServiceODataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceOData `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceOData_Kind             = "LinkedServiceOData"
	LinkedServiceOData_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceOData_Kind}.String()
	LinkedServiceOData_KindAPIVersion   = LinkedServiceOData_Kind + "." + CRDGroupVersion.String()
	LinkedServiceOData_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceOData_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceOData{}, &LinkedServiceODataList{})
}
