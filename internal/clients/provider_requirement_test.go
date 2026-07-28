package clients

import (
	"strings"
	"testing"

	accesskeyv1alpha1 "github.com/NitriKx/provider-cloudinaryprovisioning/apis/namespaced/accesskey/v1alpha1"
	policyv1alpha1 "github.com/NitriKx/provider-cloudinaryprovisioning/apis/namespaced/policy/v1alpha1"
)

const (
	testProviderSource  = "nitrikx/cloudinary-provisioning"
	testProviderVersion = "0.1.0"
)

// derivedLocalName replicates upjet's fallback derivation of the provider local
// name (pkg/terraform/files.go): the last path segment of the source.
func derivedLocalName(source string) string {
	segments := strings.Split(source, "/")
	return segments[len(segments)-1]
}

// Terraform resolves a resource to a provider through the resource type prefix,
// so every generated resource type must be served by the local name declared in
// required_providers.
func TestProviderRequirementLocalNameCoversResourceTypes(t *testing.T) {
	req := providerRequirement(testProviderSource, testProviderVersion)

	for _, resourceType := range []string{
		(&accesskeyv1alpha1.AccessKey{}).GetTerraformResourceType(),
		(&policyv1alpha1.CustomPolicy{}).GetTerraformResourceType(),
	} {
		if !strings.HasPrefix(resourceType, req.LocalName+"_") {
			t.Errorf("resource type %q is not served by provider local name %q", resourceType, req.LocalName)
		}
	}
}

// Guards the regression: relying on upjet's default derivation yields a local
// name that no resource type matches, which makes Terraform fall back to
// registry.terraform.io/hashicorp/<prefix> and fail terraform init.
func TestProviderRequirementOverridesDerivedLocalName(t *testing.T) {
	req := providerRequirement(testProviderSource, testProviderVersion)

	if req.LocalName == "" {
		t.Fatal("LocalName must be set explicitly, otherwise upjet derives it from Source")
	}
	if derived := derivedLocalName(testProviderSource); req.LocalName == derived {
		t.Fatalf("LocalName %q unexpectedly equals the derived name; this test no longer guards anything", req.LocalName)
	}
}

func TestProviderRequirementKeepsSourceAndVersion(t *testing.T) {
	req := providerRequirement(testProviderSource, testProviderVersion)

	if req.Source != testProviderSource {
		t.Errorf("Source = %q, want %q", req.Source, testProviderSource)
	}
	if req.Version != testProviderVersion {
		t.Errorf("Version = %q, want %q", req.Version, testProviderVersion)
	}
}
