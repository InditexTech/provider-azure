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

package authorization

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/config/common"
)

// Configure configures authorization group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_role_assignment", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"role_definition_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_resource_group_policy_assignment", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "resource_group_id")
	})
}
