package iam

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the cloudinaryprovisioning_principal_role_assignment resource for the namespaced provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudinaryprovisioning_principal_role_assignment", func(r *ujconfig.Resource) {
		r.Kind = "PrincipalRoleAssignment"
		r.ShortGroup = "iam"
		r.Version = "v1alpha1"
	})
}
