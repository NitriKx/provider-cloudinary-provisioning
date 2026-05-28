package accesskey

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the cloudinaryprovisioning_access_key resource.
// The api_key and api_secret fields are exported as connection details so
// compositions can consume them as Kubernetes Secrets for the upload SDK.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudinaryprovisioning_access_key", func(r *ujconfig.Resource) {
		r.Kind = "AccessKey"
		r.ShortGroup = "accesskey"
		r.Version = "v1alpha1"
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			out := map[string][]byte{}
			if v, ok := attr["api_key"].(string); ok {
				out["api_key"] = []byte(v)
			}
			if v, ok := attr["api_secret"].(string); ok {
				out["api_secret"] = []byte(v)
			}
			return out, nil
		}
	})
}
