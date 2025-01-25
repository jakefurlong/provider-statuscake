/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/jakefurlong/provider-github/config/contact_group"
	"github.com/jakefurlong/provider-github/config/heartbeat_check"
	"github.com/jakefurlong/provider-github/config/maintenance_window"
	"github.com/jakefurlong/provider-github/config/pagespeed_check"
	"github.com/jakefurlong/provider-github/config/ssl_check"
	"github.com/jakefurlong/provider-github/config/uptime_check"
)

const (
	resourcePrefix = "statuscake"
	modulePath     = "github.com/jakefurlong/provider-statuscake"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		contact_group.Configure,
		heartbeat_check.Configure,
		maintenance_window.Configure,
		pagespeed_check.Configure,
		ssl_check.Configure,
		uptime_check.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
