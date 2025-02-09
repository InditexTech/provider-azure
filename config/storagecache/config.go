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

package storagecache

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures storagesync group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_hpc_cache_nfs_target", func(r *config.Resource) {
		delete(r.References, "target_host_name")
	})
}
