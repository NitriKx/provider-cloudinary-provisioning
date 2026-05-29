package policy

import (
	"os"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func newTestProvider(t *testing.T) *ujconfig.Provider {
	t.Helper()
	schema, err := os.ReadFile("../../schema.json")
	if err != nil {
		t.Fatalf("cannot read schema.json: %v", err)
	}
	meta, err := os.ReadFile("../../provider-metadata.yaml")
	if err != nil {
		t.Fatalf("cannot read provider-metadata.yaml: %v", err)
	}
	return ujconfig.NewProvider(schema, "cloudinaryprovisioning",
		"github.com/NitriKx/provider-cloudinaryprovisioning", meta)
}

func TestConfigure_SetsExpectedFields(t *testing.T) {
	p := newTestProvider(t)
	Configure(p)
	p.ConfigureResources()

	r, ok := p.Resources["cloudinaryprovisioning_custom_policy"]
	if !ok {
		t.Fatal("cloudinaryprovisioning_custom_policy not found in provider resources after Configure")
	}
	if r.Kind != "CustomPolicy" {
		t.Errorf("Kind = %q, want %q", r.Kind, "CustomPolicy")
	}
	if r.ShortGroup != "policy" {
		t.Errorf("ShortGroup = %q, want %q", r.ShortGroup, "policy")
	}
	if r.Version != "v1alpha1" {
		t.Errorf("Version = %q, want %q", r.Version, "v1alpha1")
	}
}
