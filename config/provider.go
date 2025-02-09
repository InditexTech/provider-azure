/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	"context"
	_ "embed"

	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/xpprovider"
	"github.com/pkg/errors"

	tjconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"

	"github.com/upbound/provider-azure/config/alertsmanagement"
	"github.com/upbound/provider-azure/config/apimanagement"
	"github.com/upbound/provider-azure/config/appplatform"
	"github.com/upbound/provider-azure/config/authorization"
	"github.com/upbound/provider-azure/config/automation"
	"github.com/upbound/provider-azure/config/base"
	"github.com/upbound/provider-azure/config/botservice"
	"github.com/upbound/provider-azure/config/cache"
	"github.com/upbound/provider-azure/config/cdn"
	"github.com/upbound/provider-azure/config/certificateregistration"
	"github.com/upbound/provider-azure/config/common"
	"github.com/upbound/provider-azure/config/compute"
	"github.com/upbound/provider-azure/config/consumption"
	"github.com/upbound/provider-azure/config/containerapp"
	"github.com/upbound/provider-azure/config/containerregistry"
	"github.com/upbound/provider-azure/config/containerservice"
	"github.com/upbound/provider-azure/config/cosmosdb"
	"github.com/upbound/provider-azure/config/costmanagement"
	"github.com/upbound/provider-azure/config/datafactory"
	"github.com/upbound/provider-azure/config/dataprotection"
	"github.com/upbound/provider-azure/config/datashare"
	"github.com/upbound/provider-azure/config/dbformysql"
	"github.com/upbound/provider-azure/config/devices"
	"github.com/upbound/provider-azure/config/eventhub"
	"github.com/upbound/provider-azure/config/healthcareapis"
	"github.com/upbound/provider-azure/config/insights"
	"github.com/upbound/provider-azure/config/keyvault"
	"github.com/upbound/provider-azure/config/kusto"
	"github.com/upbound/provider-azure/config/logic"
	"github.com/upbound/provider-azure/config/management"
	"github.com/upbound/provider-azure/config/mariadb"
	"github.com/upbound/provider-azure/config/media"
	"github.com/upbound/provider-azure/config/netapp"
	"github.com/upbound/provider-azure/config/network"
	"github.com/upbound/provider-azure/config/notificationhubs"
	"github.com/upbound/provider-azure/config/operationalinsights"
	"github.com/upbound/provider-azure/config/orbital"
	"github.com/upbound/provider-azure/config/postgresql"
	"github.com/upbound/provider-azure/config/relay"
	"github.com/upbound/provider-azure/config/resource"
	"github.com/upbound/provider-azure/config/resources"
	"github.com/upbound/provider-azure/config/security"
	"github.com/upbound/provider-azure/config/servicebus"
	"github.com/upbound/provider-azure/config/sql"
	"github.com/upbound/provider-azure/config/storage"
	"github.com/upbound/provider-azure/config/storagecache"
	"github.com/upbound/provider-azure/config/storagesync"
	"github.com/upbound/provider-azure/config/streamanalytics"
	"github.com/upbound/provider-azure/config/web"
	"github.com/upbound/provider-azure/hack"
)

const (
	resourcePrefix = "azure"
	modulePath     = "github.com/upbound/provider-azure"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = []string{
	"azurerm_mssql_server_extended_auditing_policy",
	// group prefix collision
	"azurerm_api_management_group",
	"azurerm_api_management_product_group",
	"azurerm_dedicated_host_group",
	"azurerm_storage_disks_pool",
	"azurerm_storage_sync_group",
	"azurerm_storage_sync_cloud_endpoint", // depends on azurerm_storage_sync_group
	"azurerm_virtual_desktop_application_group",
	// associated with non-generated
	"azurerm_virtual_desktop_workspace_application_group_association",
	// generated name too long
	"azurerm_network_interface_application_gateway_backend_address_pool_association",
	"azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection",
	// deprecated
	"azurerm_virtual_machine_scale_set",
	"azurerm_virtual_machine_configuration_policy_assignment",
	"azurerm_virtual_machine_scale_set_extension",
	"azurerm_sql_server",
	"azurerm_video_analyzer",
	"azurerm_video_analyzer_edge_module",
	"azurerm_sql_managed_database",
	"azurerm_sql_managed_instance",
	"azurerm_sql_managed_instance_active_directory_administrator",
	"azurerm_sql_managed_instance_failover_group",
	"azurerm_sql_active_directory_administrator",
	"azurerm_sql_database",
	"azurerm_sql_elasticpool",
	"azurerm_sql_firewall_rule",
	"azurerm_site_recovery_replicated_vm", // depends on azurerm_virtual_machine
	// irrelevant
	"azurerm_virtual_desktop_application",
	"azurerm_virtual_desktop_host_pool",
	"azurerm_virtual_desktop_workspace",
	"azurerm_virtual_desktop_scaling_plan",                // depends on azurerm_virtual_desktop_host_pool
	"azurerm_virtual_desktop_host_pool_registration_info", // depends on azurerm_virtual_desktop_host_pool
	// other upjet issues
	"azurerm_container_registry_task",
	"azurerm_dashboard",
	// doc not found in Terraform Azurerm provider
	"azurerm_virtual_network_dns_servers",
	// unsupported sensitive field type
	"azurerm_security_center_automation",
	"azurerm_data_factory_trigger_tumbling_window",
	"azurerm_storage_share_file",
	"azurerm_sql_virtual_network_rule",
	"azurerm_virtual_desktop_workspace",
	"azurerm_data_lake_analytics_account",
	"azurerm_log_analytics_storage_insights",
	"azurerm_virtual_hub_bgp_connection",
	"azurerm_automation_dsc_configuration",
	"azurerm_automation_dsc_nodeconfiguration", // depends on azurerm_automation_dsc_configuration
	"azurerm_monitor_log_profile",
	"azurerm_machine_learning_inference_cluster",
	"azurerm_sql_failover_group",
	"azurerm_logic_app_integration_account_certificate",
	"azurerm_container_group",
	// Azure are officially halting the preview of Azure Disk Pools, and it will not be made generally available.
	"azurerm_disk_pool_iscsi_target",
	// DEPRECATED: azurerm_service_plan should be used instead
	"azurerm_app_service_source_control_token",
	// DEPRECATED:
	"azurerm_cdn_frontdoor_route_disable_link_to_default_domain",
}

// workaround for the TF Azure v3.57.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(ctx context.Context, generationProvider bool) (*tjconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p, err = xpprovider.GetProviderSchema(ctx)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		tjconfig.WithShortName("azure"),
		tjconfig.WithRootGroup("azure.upbound.io"),
		tjconfig.WithIncludeList(CLIReconciledResourceList()),
		tjconfig.WithNoForkIncludeList(NoForkResourceList()),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithDefaultResourceOptions(ResourceConfigurator()),
		tjconfig.WithReferenceInjectors([]tjconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		tjconfig.WithFeaturesPackage("internal/features"),
		tjconfig.WithMainTemplate(hack.MainTemplate),
		tjconfig.WithTerraformProvider(p),
	)
	// "azure" group contains resources that actually do not have a specific
	// group, e.g. ResourceGroup with APIVersion "azure.upbound.io/v1beta1".
	// We need to include the controllers for this group into the base packages
	// list to get their controllers packaged together with the config package
	// controllers (provider family config package).
	for _, c := range []string{"internal/controller/azure/resourcegroup", "internal/controller/azure/resourceproviderregistration", "internal/controller/azure/subscription"} {
		pc.BasePackages.ControllerMap[c] = "config"
	}

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		authorization.Configure,
		alertsmanagement.Configure,
		network.Configure,
		management.Configure,
		media.Configure,
		cache.Configure,
		resource.Configure,
		resources.Configure,
		containerapp.Configure,
		containerservice.Configure,
		postgresql.Configure,
		cosmosdb.Configure,
		sql.Configure,
		storage.Configure,
		operationalinsights.Configure,
		insights.Configure,
		devices.Configure,
		datafactory.Configure,
		apimanagement.Configure,
		healthcareapis.Configure,
		logic.Configure,
		security.Configure,
		base.Configure,
		botservice.Configure,
		datashare.Configure,
		notificationhubs.Configure,
		storagesync.Configure,
		keyvault.Configure,
		eventhub.Configure,
		mariadb.Configure,
		compute.Configure,
		containerregistry.Configure,
		dbformysql.Configure,
		netapp.Configure,
		dataprotection.Configure,
		kusto.Configure,
		storagecache.Configure,
		servicebus.Configure,
		consumption.Configure,
		streamanalytics.Configure,
		costmanagement.Configure,
		automation.Configure,
		web.Configure,
		relay.Configure,
		cdn.Configure,
		certificateregistration.Configure,
		orbital.Configure,
		appplatform.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()

	// This function runs after the special configurations were applied. However, if some references were configured in
	// the configuration files, the reference generator does not override them.
	for _, r := range pc.Resources {
		delete(r.References, "resource_group_name")
		if err := common.AddCommonReferences(r); err != nil {
			panic(err)
		}
	}
	return pc, nil
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// NoForkResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func NoForkResourceList() []string {
	l := make([]string, len(NoForkExternalNameConfigs))
	i := 0
	for name := range NoForkExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
