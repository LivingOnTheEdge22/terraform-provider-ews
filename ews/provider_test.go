package ews

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccProviderConfigure sync.Once

func ThreeValidPoPs() []string {
	validPoPs := []string{"hkg", "lon", "iad"}
	if v := os.Getenv("EWS_BASE_URL"); v == "https://my.impervaservices.com/api/prov/v1" {
		validPoPs[0] = "sus"
		validPoPs[1] = "bst"
		validPoPs[2] = "ogn"
	}
	return validPoPs
}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"incapsula": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	testAccProviderConfigure.Do(func() {
		if v := os.Getenv("INCAPSULA_API_ID"); v == "" {
			t.Fatal("INCAPSULA_API_ID must be set for acceptance tests")
		}

		if v := os.Getenv("INCAPSULA_API_KEY"); v == "" {
			t.Fatal("INCAPSULA_API_KEY must be set for acceptance tests")
		}

		err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
		if err != nil {
			t.Fatal(err)
		}
	})
}
