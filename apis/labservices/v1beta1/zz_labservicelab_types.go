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

type AdminUserObservation struct {

	// The username to use when signing in to Lab Service Lab VMs.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AdminUserParameters struct {

	// The password for the user.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username to use when signing in to Lab Service Lab VMs.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type AutoShutdownObservation struct {

	// The amount of time a VM will stay running after a user disconnects if this behavior is enabled. This value must be formatted as an ISO 8601 string.
	DisconnectDelay *string `json:"disconnectDelay,omitempty" tf:"disconnect_delay,omitempty"`

	// The amount of time a VM will idle before it is shutdown if this behavior is enabled. This value must be formatted as an ISO 8601 string.
	IdleDelay *string `json:"idleDelay,omitempty" tf:"idle_delay,omitempty"`

	// The amount of time a VM will stay running before it is shutdown if no connection is made and this behavior is enabled. This value must be formatted as an ISO 8601 string.
	NoConnectDelay *string `json:"noConnectDelay,omitempty" tf:"no_connect_delay,omitempty"`

	// A VM will get shutdown when it has idled for a period of time. Possible values are LowUsage and UserAbsence.
	ShutdownOnIdle *string `json:"shutdownOnIdle,omitempty" tf:"shutdown_on_idle,omitempty"`
}

type AutoShutdownParameters struct {

	// The amount of time a VM will stay running after a user disconnects if this behavior is enabled. This value must be formatted as an ISO 8601 string.
	// +kubebuilder:validation:Optional
	DisconnectDelay *string `json:"disconnectDelay,omitempty" tf:"disconnect_delay,omitempty"`

	// The amount of time a VM will idle before it is shutdown if this behavior is enabled. This value must be formatted as an ISO 8601 string.
	// +kubebuilder:validation:Optional
	IdleDelay *string `json:"idleDelay,omitempty" tf:"idle_delay,omitempty"`

	// The amount of time a VM will stay running before it is shutdown if no connection is made and this behavior is enabled. This value must be formatted as an ISO 8601 string.
	// +kubebuilder:validation:Optional
	NoConnectDelay *string `json:"noConnectDelay,omitempty" tf:"no_connect_delay,omitempty"`

	// A VM will get shutdown when it has idled for a period of time. Possible values are LowUsage and UserAbsence.
	// +kubebuilder:validation:Optional
	ShutdownOnIdle *string `json:"shutdownOnIdle,omitempty" tf:"shutdown_on_idle,omitempty"`
}

type ConnectionSettingObservation struct {

	// The enabled access level for Client Access over RDP. Possible value is Public.
	ClientRdpAccess *string `json:"clientRdpAccess,omitempty" tf:"client_rdp_access,omitempty"`

	// The enabled access level for Client Access over SSH. Possible value is Public.
	ClientSSHAccess *string `json:"clientSshAccess,omitempty" tf:"client_ssh_access,omitempty"`
}

type ConnectionSettingParameters struct {

	// The enabled access level for Client Access over RDP. Possible value is Public.
	// +kubebuilder:validation:Optional
	ClientRdpAccess *string `json:"clientRdpAccess,omitempty" tf:"client_rdp_access,omitempty"`

	// The enabled access level for Client Access over SSH. Possible value is Public.
	// +kubebuilder:validation:Optional
	ClientSSHAccess *string `json:"clientSshAccess,omitempty" tf:"client_ssh_access,omitempty"`
}

type ImageReferenceObservation struct {

	// The resource ID of the image. Changing this forces a new resource to be created.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image offer if applicable. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The image publisher. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// A sku block as defined below.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The image version specified on creation. Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ImageReferenceParameters struct {

	// The resource ID of the image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image offer if applicable. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The image publisher. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// A sku block as defined below.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The image version specified on creation. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LabServiceLabObservation struct {

	// An auto_shutdown block as defined below.
	AutoShutdown []AutoShutdownObservation `json:"autoShutdown,omitempty" tf:"auto_shutdown,omitempty"`

	// A connection_setting block as defined below.
	ConnectionSetting []ConnectionSettingObservation `json:"connectionSetting,omitempty" tf:"connection_setting,omitempty"`

	// The description of the Lab Service Lab.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Lab Service Lab.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource ID of the Lab Plan that is used during resource creation to provide defaults and acts as a permission container when creating a Lab Service Lab via labs.azure.com.
	LabPlanID *string `json:"labPlanId,omitempty" tf:"lab_plan_id,omitempty"`

	// The Azure Region where the Lab Service Lab should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A network block as defined below.
	Network []NetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// The name of the Resource Group where the Lab Service Lab should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A roster block as defined below.
	Roster []RosterObservation `json:"roster,omitempty" tf:"roster,omitempty"`

	// A security block as defined below.
	Security []SecurityObservation `json:"security,omitempty" tf:"security,omitempty"`

	// A mapping of tags which should be assigned to the Lab Service Lab.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The title of the Lab Service Lab.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// A virtual_machine block as defined below.
	VirtualMachine []VirtualMachineObservation `json:"virtualMachine,omitempty" tf:"virtual_machine,omitempty"`
}

type LabServiceLabParameters struct {

	// An auto_shutdown block as defined below.
	// +kubebuilder:validation:Optional
	AutoShutdown []AutoShutdownParameters `json:"autoShutdown,omitempty" tf:"auto_shutdown,omitempty"`

	// A connection_setting block as defined below.
	// +kubebuilder:validation:Optional
	ConnectionSetting []ConnectionSettingParameters `json:"connectionSetting,omitempty" tf:"connection_setting,omitempty"`

	// The description of the Lab Service Lab.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID of the Lab Plan that is used during resource creation to provide defaults and acts as a permission container when creating a Lab Service Lab via labs.azure.com.
	// +kubebuilder:validation:Optional
	LabPlanID *string `json:"labPlanId,omitempty" tf:"lab_plan_id,omitempty"`

	// The Azure Region where the Lab Service Lab should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A network block as defined below.
	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// The name of the Resource Group where the Lab Service Lab should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A roster block as defined below.
	// +kubebuilder:validation:Optional
	Roster []RosterParameters `json:"roster,omitempty" tf:"roster,omitempty"`

	// A security block as defined below.
	// +kubebuilder:validation:Optional
	Security []SecurityParameters `json:"security,omitempty" tf:"security,omitempty"`

	// A mapping of tags which should be assigned to the Lab Service Lab.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The title of the Lab Service Lab.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// A virtual_machine block as defined below.
	// +kubebuilder:validation:Optional
	VirtualMachine []VirtualMachineParameters `json:"virtualMachine,omitempty" tf:"virtual_machine,omitempty"`
}

type NetworkObservation struct {

	// The resource ID of the Load Balancer for the network profile of the Lab Service Lab.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// The resource ID of the Public IP for the network profile of the Lab Service Lab.
	PublicIPID *string `json:"publicIpId,omitempty" tf:"public_ip_id,omitempty"`

	// The resource ID of the Subnet for the network profile of the Lab Service Lab.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkParameters struct {

	// The resource ID of the Subnet for the network profile of the Lab Service Lab.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type NonAdminUserObservation struct {

	// The username to use when signing in to Lab Service Lab VMs.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type NonAdminUserParameters struct {

	// The password for the user.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username to use when signing in to Lab Service Lab VMs.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type RosterObservation struct {

	// The AAD group ID which this Lab Service Lab roster is populated from.
	ActiveDirectoryGroupID *string `json:"activeDirectoryGroupId,omitempty" tf:"active_directory_group_id,omitempty"`

	// The base URI identifying the lms instance.
	LmsInstance *string `json:"lmsInstance,omitempty" tf:"lms_instance,omitempty"`

	// The unique id of the Azure Lab Service tool in the lms.
	LtiClientID *string `json:"ltiClientId,omitempty" tf:"lti_client_id,omitempty"`

	// The unique context identifier for the Lab Service Lab in the lms.
	LtiContextID *string `json:"ltiContextId,omitempty" tf:"lti_context_id,omitempty"`

	// The URI of the names and roles service endpoint on the lms for the class attached to this Lab Service Lab.
	LtiRosterEndpoint *string `json:"ltiRosterEndpoint,omitempty" tf:"lti_roster_endpoint,omitempty"`
}

type RosterParameters struct {

	// The AAD group ID which this Lab Service Lab roster is populated from.
	// +kubebuilder:validation:Optional
	ActiveDirectoryGroupID *string `json:"activeDirectoryGroupId,omitempty" tf:"active_directory_group_id,omitempty"`

	// The base URI identifying the lms instance.
	// +kubebuilder:validation:Optional
	LmsInstance *string `json:"lmsInstance,omitempty" tf:"lms_instance,omitempty"`

	// The unique id of the Azure Lab Service tool in the lms.
	// +kubebuilder:validation:Optional
	LtiClientID *string `json:"ltiClientId,omitempty" tf:"lti_client_id,omitempty"`

	// The unique context identifier for the Lab Service Lab in the lms.
	// +kubebuilder:validation:Optional
	LtiContextID *string `json:"ltiContextId,omitempty" tf:"lti_context_id,omitempty"`

	// The URI of the names and roles service endpoint on the lms for the class attached to this Lab Service Lab.
	// +kubebuilder:validation:Optional
	LtiRosterEndpoint *string `json:"ltiRosterEndpoint,omitempty" tf:"lti_roster_endpoint,omitempty"`
}

type SecurityObservation struct {

	// Is open access enabled to allow any user or only specified users to register to a Lab Service Lab?
	OpenAccessEnabled *bool `json:"openAccessEnabled,omitempty" tf:"open_access_enabled,omitempty"`

	// The registration code for the Lab Service Lab.
	RegistrationCode *string `json:"registrationCode,omitempty" tf:"registration_code,omitempty"`
}

type SecurityParameters struct {

	// Is open access enabled to allow any user or only specified users to register to a Lab Service Lab?
	// +kubebuilder:validation:Required
	OpenAccessEnabled *bool `json:"openAccessEnabled" tf:"open_access_enabled,omitempty"`
}

type SkuObservation struct {

	// The capacity for the SKU. Possible values are between 0 and 400.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the SKU. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuParameters struct {

	// The capacity for the SKU. Possible values are between 0 and 400.
	// +kubebuilder:validation:Required
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// The name of the SKU. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type VirtualMachineObservation struct {

	// Is flagged to pre-install dedicated GPU drivers? Defaults to false. Changing this forces a new resource to be created.
	AdditionalCapabilityGpuDriversInstalled *bool `json:"additionalCapabilityGpuDriversInstalled,omitempty" tf:"additional_capability_gpu_drivers_installed,omitempty"`

	// An admin_user block as defined below.
	AdminUser []AdminUserObservation `json:"adminUser,omitempty" tf:"admin_user,omitempty"`

	// The create option to indicate what Lab Service Lab VMs are created from. Possible values are Image and TemplateVM. Defaults to Image. Changing this forces a new resource to be created.
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// An image_reference block as defined below.
	ImageReference []ImageReferenceObservation `json:"imageReference,omitempty" tf:"image_reference,omitempty"`

	// A non_admin_user block as defined below.
	NonAdminUser []NonAdminUserObservation `json:"nonAdminUser,omitempty" tf:"non_admin_user,omitempty"`

	// Is the shared password enabled with the same password for all user VMs? Defaults to false. Changing this forces a new resource to be created.
	SharedPasswordEnabled *bool `json:"sharedPasswordEnabled,omitempty" tf:"shared_password_enabled,omitempty"`

	// A sku block as defined below.
	Sku []SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// The initial quota allocated to each Lab Service Lab user. Defaults to PT0S. This value must be formatted as an ISO 8601 string.
	UsageQuota *string `json:"usageQuota,omitempty" tf:"usage_quota,omitempty"`
}

type VirtualMachineParameters struct {

	// Is flagged to pre-install dedicated GPU drivers? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AdditionalCapabilityGpuDriversInstalled *bool `json:"additionalCapabilityGpuDriversInstalled,omitempty" tf:"additional_capability_gpu_drivers_installed,omitempty"`

	// An admin_user block as defined below.
	// +kubebuilder:validation:Required
	AdminUser []AdminUserParameters `json:"adminUser" tf:"admin_user,omitempty"`

	// The create option to indicate what Lab Service Lab VMs are created from. Possible values are Image and TemplateVM. Defaults to Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// An image_reference block as defined below.
	// +kubebuilder:validation:Required
	ImageReference []ImageReferenceParameters `json:"imageReference" tf:"image_reference,omitempty"`

	// A non_admin_user block as defined below.
	// +kubebuilder:validation:Optional
	NonAdminUser []NonAdminUserParameters `json:"nonAdminUser,omitempty" tf:"non_admin_user,omitempty"`

	// Is the shared password enabled with the same password for all user VMs? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SharedPasswordEnabled *bool `json:"sharedPasswordEnabled,omitempty" tf:"shared_password_enabled,omitempty"`

	// A sku block as defined below.
	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// The initial quota allocated to each Lab Service Lab user. Defaults to PT0S. This value must be formatted as an ISO 8601 string.
	// +kubebuilder:validation:Optional
	UsageQuota *string `json:"usageQuota,omitempty" tf:"usage_quota,omitempty"`
}

// LabServiceLabSpec defines the desired state of LabServiceLab
type LabServiceLabSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LabServiceLabParameters `json:"forProvider"`
}

// LabServiceLabStatus defines the observed state of LabServiceLab.
type LabServiceLabStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LabServiceLabObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LabServiceLab is the Schema for the LabServiceLabs API. Manages a Lab Service Lab.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LabServiceLab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.security)",message="security is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.title)",message="title is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.virtualMachine)",message="virtualMachine is a required parameter"
	Spec   LabServiceLabSpec   `json:"spec"`
	Status LabServiceLabStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LabServiceLabList contains a list of LabServiceLabs
type LabServiceLabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LabServiceLab `json:"items"`
}

// Repository type metadata.
var (
	LabServiceLab_Kind             = "LabServiceLab"
	LabServiceLab_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LabServiceLab_Kind}.String()
	LabServiceLab_KindAPIVersion   = LabServiceLab_Kind + "." + CRDGroupVersion.String()
	LabServiceLab_GroupVersionKind = CRDGroupVersion.WithKind(LabServiceLab_Kind)
)

func init() {
	SchemeBuilder.Register(&LabServiceLab{}, &LabServiceLabList{})
}
