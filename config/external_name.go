package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// All three resources use IdentifierFromProvider because Cloudinary
// generates the IDs server-side on create; users never specify them.
// cloudinaryprovisioning_product_environment is intentionally excluded from
// v1 — it is managed manually in the Cloudinary console.
var ExternalNameConfigs = map[string]config.ExternalName{
	// TF state ID is the composite "{product_environment_id}/{api_key}".
	"cloudinaryprovisioning_access_key": config.IdentifierFromProvider,
	// TF computed id field.
	"cloudinaryprovisioning_custom_policy": config.IdentifierFromProvider,
	// TF state ID is a 5-tuple composite.
	"cloudinaryprovisioning_principal_role_assignment": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
