package policy

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the cloudinaryprovisioning_custom_policy resource for the namespaced provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudinaryprovisioning_custom_policy", func(r *ujconfig.Resource) {
		r.Kind = "CustomPolicy"
		r.ShortGroup = "policy"
		r.Version = "v1alpha1"
	})
}
