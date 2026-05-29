package accesskey

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

	r, ok := p.Resources["cloudinaryprovisioning_access_key"]
	if !ok {
		t.Fatal("cloudinaryprovisioning_access_key not found in provider resources after Configure")
	}
	if r.Kind != "AccessKey" {
		t.Errorf("Kind = %q, want %q", r.Kind, "AccessKey")
	}
	if r.ShortGroup != "accesskey" {
		t.Errorf("ShortGroup = %q, want %q", r.ShortGroup, "accesskey")
	}
	if r.Version != "v1alpha1" {
		t.Errorf("Version = %q, want %q", r.Version, "v1alpha1")
	}
}

func TestConfigure_ConnectionDetailsFn_ExportsCredentials(t *testing.T) {
	p := newTestProvider(t)
	Configure(p)
	p.ConfigureResources()

	r := p.Resources["cloudinaryprovisioning_access_key"]
	if r.Sensitive.AdditionalConnectionDetailsFn == nil {
		t.Fatal("AdditionalConnectionDetailsFn is nil")
	}

	out, err := r.Sensitive.AdditionalConnectionDetailsFn(map[string]any{
		"api_key":    "test-key",
		"api_secret": "test-secret",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(out["api_key"]) != "test-key" {
		t.Errorf("api_key = %q, want %q", out["api_key"], "test-key")
	}
	if string(out["api_secret"]) != "test-secret" {
		t.Errorf("api_secret = %q, want %q", out["api_secret"], "test-secret")
	}
}

func TestConfigure_ConnectionDetailsFn_MissingFieldsProduceNoError(t *testing.T) {
	p := newTestProvider(t)
	Configure(p)
	p.ConfigureResources()

	r := p.Resources["cloudinaryprovisioning_access_key"]
	out, err := r.Sensitive.AdditionalConnectionDetailsFn(map[string]any{})
	if err != nil {
		t.Fatalf("unexpected error on empty attr map: %v", err)
	}
	if len(out) != 0 {
		t.Errorf("expected empty output for empty attr map, got %v", out)
	}
}
