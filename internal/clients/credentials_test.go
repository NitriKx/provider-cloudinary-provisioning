package clients

import (
	"testing"
)

// applyProvisioningCredentials replicates the credential-mapping logic from
// TerraformSetupBuilder for the Provisioning provider.
func applyProvisioningCredentials(creds map[string]string) map[string]any {
	cfg := map[string]any{}
	if v := creds["account_url"]; v != "" {
		cfg["account_url"] = v
	} else {
		if v := creds["account_id"]; v != "" {
			cfg["account_id"] = v
		}
		if v := creds["api_key"]; v != "" {
			cfg["api_key"] = v
		}
		if v := creds["api_secret"]; v != "" {
			cfg["api_secret"] = v
		}
	}
	return cfg
}

func TestProvisioningCredentials_AccountURLTakesPrecedence(t *testing.T) {
	creds := map[string]string{
		"account_url": "account://key:secret@accountid",
		"account_id":  "ignored",
		"api_key":     "ignored",
		"api_secret":  "ignored",
	}
	cfg := applyProvisioningCredentials(creds)

	if v, ok := cfg["account_url"]; !ok || v != "account://key:secret@accountid" {
		t.Errorf("account_url = %v, want %q", v, "account://key:secret@accountid")
	}
	if _, ok := cfg["account_id"]; ok {
		t.Error("account_id should be absent when account_url is set")
	}
}

func TestProvisioningCredentials_TriplePath(t *testing.T) {
	creds := map[string]string{
		"account_id": "myaccount",
		"api_key":    "mykey",
		"api_secret": "mysecret",
	}
	cfg := applyProvisioningCredentials(creds)

	if _, ok := cfg["account_url"]; ok {
		t.Error("account_url should be absent when triple path is used")
	}
	if v, ok := cfg["account_id"]; !ok || v != "myaccount" {
		t.Errorf("account_id = %v, want %q", v, "myaccount")
	}
	if v, ok := cfg["api_key"]; !ok || v != "mykey" {
		t.Errorf("api_key = %v, want %q", v, "mykey")
	}
	if v, ok := cfg["api_secret"]; !ok || v != "mysecret" {
		t.Errorf("api_secret = %v, want %q", v, "mysecret")
	}
}

func TestProvisioningCredentials_EmptyCredsProducesEmptyConfig(t *testing.T) {
	cfg := applyProvisioningCredentials(map[string]string{})
	if len(cfg) != 0 {
		t.Errorf("expected empty config for empty credentials, got %v", cfg)
	}
}
