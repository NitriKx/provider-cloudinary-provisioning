package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	accesskeyCluster "github.com/NitriKx/provider-cloudinaryprovisioning/config/cluster/accesskey"
	iamCluster "github.com/NitriKx/provider-cloudinaryprovisioning/config/cluster/iam"
	policyCluster "github.com/NitriKx/provider-cloudinaryprovisioning/config/cluster/policy"
	accesskeyNamespaced "github.com/NitriKx/provider-cloudinaryprovisioning/config/namespaced/accesskey"
	iamNamespaced "github.com/NitriKx/provider-cloudinaryprovisioning/config/namespaced/iam"
	policyNamespaced "github.com/NitriKx/provider-cloudinaryprovisioning/config/namespaced/policy"
)

const (
	resourcePrefix = "cloudinaryprovisioning"
	modulePath     = "github.com/NitriKx/provider-cloudinaryprovisioning"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudinaryprovisioning.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		accesskeyCluster.Configure,
		policyCluster.Configure,
		iamCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudinaryprovisioning.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		accesskeyNamespaced.Configure,
		policyNamespaced.Configure,
		iamNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
