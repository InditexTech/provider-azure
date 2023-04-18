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

type ApplicationInsightsObservation struct {

	// The App ID associated with this Application Insights component.
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Specifies the type of Application Insights to create. Valid values are ios for iOS, java for Java web, MobileCenter for App Center, Node.JS for Node.js, other for General, phone for Windows Phone, store for Windows Store and web for ASP.NET. Please note these values are case sensitive; unmatched values are treated as ASP.NET by Azure. Changing this forces a new resource to be created.
	ApplicationType *string `json:"applicationType,omitempty" tf:"application_type,omitempty"`

	// Specifies the Application Insights component daily data volume cap in GB.
	DailyDataCapInGb *float64 `json:"dailyDataCapInGb,omitempty" tf:"daily_data_cap_in_gb,omitempty"`

	// Specifies if a notification email will be send when the daily data volume cap is met.
	DailyDataCapNotificationsDisabled *bool `json:"dailyDataCapNotificationsDisabled,omitempty" tf:"daily_data_cap_notifications_disabled,omitempty"`

	// By default the real client IP is masked as 0.0.0.0 in the logs. Use this argument to disable masking and log the real client IP. Defaults to false.
	DisableIPMasking *bool `json:"disableIpMasking,omitempty" tf:"disable_ip_masking,omitempty"`

	// Should the Application Insights component force users to create their own storage account for profiling? Defaults to false.
	ForceCustomerStorageForProfiler *bool `json:"forceCustomerStorageForProfiler,omitempty" tf:"force_customer_storage_for_profiler,omitempty"`

	// The ID of the Application Insights component.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Should the Application Insights component support ingestion over the Public Internet? Defaults to true.
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// Should the Application Insights component support querying over the Public Internet? Defaults to true.
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	// Disable Non-Azure AD based Auth. Defaults to false.
	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Application Insights component. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the retention period in days. Possible values are 30, 60, 90, 120, 180, 270, 365, 550 or 730. Defaults to 90.
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry. Defaults to 100.
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of a log analytics workspace resource. Changing this forces a new resource to be created.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type ApplicationInsightsParameters struct {

	// Specifies the type of Application Insights to create. Valid values are ios for iOS, java for Java web, MobileCenter for App Center, Node.JS for Node.js, other for General, phone for Windows Phone, store for Windows Store and web for ASP.NET. Please note these values are case sensitive; unmatched values are treated as ASP.NET by Azure. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ApplicationType *string `json:"applicationType,omitempty" tf:"application_type,omitempty"`

	// Specifies the Application Insights component daily data volume cap in GB.
	// +kubebuilder:validation:Optional
	DailyDataCapInGb *float64 `json:"dailyDataCapInGb,omitempty" tf:"daily_data_cap_in_gb,omitempty"`

	// Specifies if a notification email will be send when the daily data volume cap is met.
	// +kubebuilder:validation:Optional
	DailyDataCapNotificationsDisabled *bool `json:"dailyDataCapNotificationsDisabled,omitempty" tf:"daily_data_cap_notifications_disabled,omitempty"`

	// By default the real client IP is masked as 0.0.0.0 in the logs. Use this argument to disable masking and log the real client IP. Defaults to false.
	// +kubebuilder:validation:Optional
	DisableIPMasking *bool `json:"disableIpMasking,omitempty" tf:"disable_ip_masking,omitempty"`

	// Should the Application Insights component force users to create their own storage account for profiling? Defaults to false.
	// +kubebuilder:validation:Optional
	ForceCustomerStorageForProfiler *bool `json:"forceCustomerStorageForProfiler,omitempty" tf:"force_customer_storage_for_profiler,omitempty"`

	// Should the Application Insights component support ingestion over the Public Internet? Defaults to true.
	// +kubebuilder:validation:Optional
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// Should the Application Insights component support querying over the Public Internet? Defaults to true.
	// +kubebuilder:validation:Optional
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	// Disable Non-Azure AD based Auth. Defaults to false.
	// +kubebuilder:validation:Optional
	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Application Insights component. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the retention period in days. Possible values are 30, 60, 90, 120, 180, 270, 365, 550 or 730. Defaults to 90.
	// +kubebuilder:validation:Optional
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry. Defaults to 100.
	// +kubebuilder:validation:Optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of a log analytics workspace resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// ApplicationInsightsSpec defines the desired state of ApplicationInsights
type ApplicationInsightsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsParameters `json:"forProvider"`
}

// ApplicationInsightsStatus defines the observed state of ApplicationInsights.
type ApplicationInsightsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsights is the Schema for the ApplicationInsightss API. Manages an Application Insights component.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsights struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.applicationType)",message="applicationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.location)",message="location is a required parameter"
	Spec   ApplicationInsightsSpec   `json:"spec"`
	Status ApplicationInsightsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsList contains a list of ApplicationInsightss
type ApplicationInsightsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsights `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsights_Kind             = "ApplicationInsights"
	ApplicationInsights_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsights_Kind}.String()
	ApplicationInsights_KindAPIVersion   = ApplicationInsights_Kind + "." + CRDGroupVersion.String()
	ApplicationInsights_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsights_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsights{}, &ApplicationInsightsList{})
}
