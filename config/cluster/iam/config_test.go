package iam

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

	r, ok := p.Resources["cloudinaryprovisioning_principal_role_assignment"]
	if !ok {
		t.Fatal("cloudinaryprovisioning_principal_role_assignment not found in provider resources after Configure")
	}
	if r.Kind != "PrincipalRoleAssignment" {
		t.Errorf("Kind = %q, want %q", r.Kind, "PrincipalRoleAssignment")
	}
	if r.ShortGroup != "iam" {
		t.Errorf("ShortGroup = %q, want %q", r.ShortGroup, "iam")
	}
	if r.Version != "v1alpha1" {
		t.Errorf("Version = %q, want %q", r.Version, "v1alpha1")
	}
}
