package jerakia

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	JERAKIA_URL   = os.Getenv("JERAKIA_URL")
	JERAKIA_TOKEN = os.Getenv("JERAKIA_TOKEN")
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"jerakia": testAccProvider,
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	if JERAKIA_URL == "" {
		t.Fatal("JERAKIA_URL must be set for acceptance tests")
	}

	if JERAKIA_TOKEN == "" {
		t.Fatal("JERAKIA_TOKEN must be set for acceptance tests")
	}
}

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
}
