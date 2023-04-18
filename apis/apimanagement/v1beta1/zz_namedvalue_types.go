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

type NamedValueObservation struct {

	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The display name of this API Management Named Value.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the API Management Named Value.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies whether the API Management Named Value is secret. Valid values are true or false. The default value is false.
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// A list of tags to be applied to the API Management Named Value.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A value_from_key_vault block as defined below.
	ValueFromKeyVault []ValueFromKeyVaultObservation `json:"valueFromKeyVault,omitempty" tf:"value_from_key_vault,omitempty"`
}

type NamedValueParameters struct {

	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The display name of this API Management Named Value.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies whether the API Management Named Value is secret. Valid values are true or false. The default value is false.
	// +kubebuilder:validation:Optional
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// A list of tags to be applied to the API Management Named Value.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A value_from_key_vault block as defined below.
	// +kubebuilder:validation:Optional
	ValueFromKeyVault []ValueFromKeyVaultParameters `json:"valueFromKeyVault,omitempty" tf:"value_from_key_vault,omitempty"`

	// The value of this API Management Named Value.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type ValueFromKeyVaultObservation struct {

	// The client ID of User Assigned Identity, for the API Management Service, which will be used to access the key vault secret. The System Assigned Identity will be used in absence.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// The resource ID of the Key Vault Secret.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`
}

type ValueFromKeyVaultParameters struct {

	// The client ID of User Assigned Identity, for the API Management Service, which will be used to access the key vault secret. The System Assigned Identity will be used in absence.
	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// The resource ID of the Key Vault Secret.
	// +kubebuilder:validation:Required
	SecretID *string `json:"secretId" tf:"secret_id,omitempty"`
}

// NamedValueSpec defines the desired state of NamedValue
type NamedValueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamedValueParameters `json:"forProvider"`
}

// NamedValueStatus defines the observed state of NamedValue.
type NamedValueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamedValueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NamedValue is the Schema for the NamedValues API. Manages an API Management Named Value.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NamedValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   NamedValueSpec   `json:"spec"`
	Status NamedValueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamedValueList contains a list of NamedValues
type NamedValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamedValue `json:"items"`
}

// Repository type metadata.
var (
	NamedValue_Kind             = "NamedValue"
	NamedValue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NamedValue_Kind}.String()
	NamedValue_KindAPIVersion   = NamedValue_Kind + "." + CRDGroupVersion.String()
	NamedValue_GroupVersionKind = CRDGroupVersion.WithKind(NamedValue_Kind)
)

func init() {
	SchemeBuilder.Register(&NamedValue{}, &NamedValueList{})
}
