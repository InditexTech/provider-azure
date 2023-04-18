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

type AccountObservation struct {

	// Catalog endpoint.
	CatalogEndpoint *string `json:"catalogEndpoint,omitempty" tf:"catalog_endpoint,omitempty"`

	// Guardian endpoint.
	GuardianEndpoint *string `json:"guardianEndpoint,omitempty" tf:"guardian_endpoint,omitempty"`

	// The ID of the Purview Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Purview Account should exist. Changing this forces a new Purview Account to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name which should be used for the new Resource Group where Purview Account creates the managed resources. Changing this forces a new Purview Account to be created.
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// A managed_resources block as defined below.
	ManagedResources []ManagedResourcesObservation `json:"managedResources,omitempty" tf:"managed_resources,omitempty"`

	// Should the Purview Account be visible to the public network? Defaults to true.
	PublicNetworkEnabled *bool `json:"publicNetworkEnabled,omitempty" tf:"public_network_enabled,omitempty"`

	// The name of the Resource Group where the Purview Account should exist. Changing this forces a new Purview Account to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Scan endpoint.
	ScanEndpoint *string `json:"scanEndpoint,omitempty" tf:"scan_endpoint,omitempty"`

	// A mapping of tags which should be assigned to the Purview Account.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountParameters struct {

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Purview Account should exist. Changing this forces a new Purview Account to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name which should be used for the new Resource Group where Purview Account creates the managed resources. Changing this forces a new Purview Account to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameRef *v1.Reference `json:"managedResourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameSelector *v1.Selector `json:"managedResourceGroupNameSelector,omitempty" tf:"-"`

	// Should the Purview Account be visible to the public network? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkEnabled *bool `json:"publicNetworkEnabled,omitempty" tf:"public_network_enabled,omitempty"`

	// The name of the Resource Group where the Purview Account should exist. Changing this forces a new Purview Account to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Purview Account.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Purview Account.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Purview Account. Possible values are UserAssigned and SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Purview Account.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Purview Account. Possible values are UserAssigned and SystemAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ManagedResourcesObservation struct {

	// The ID of the managed event hub namespace.
	EventHubNamespaceID *string `json:"eventHubNamespaceId,omitempty" tf:"event_hub_namespace_id,omitempty"`

	// The ID of the managed resource group.
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// The ID of the managed storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type ManagedResourcesParameters struct {
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API. Manages a Purview Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.identity)",message="identity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.location)",message="location is a required parameter"
	Spec   AccountSpec   `json:"spec"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
