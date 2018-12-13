package jerakia

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	fixtures "github.com/jerakia/go-jerakia/testing"
)

func TestAccDataSourceLookup_basic(t *testing.T) {
	resultJSON, err := json.Marshal(fixtures.LookupBasicResult.Payload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(resultJSON)),
				),
			},
		},
	})
}

func TestAccDataSourceLookup_singleBool(t *testing.T) {
	resultJSON, err := json.Marshal(fixtures.LookupSingleBoolResult.Payload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_singleBool,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(resultJSON)),
				),
			},
		},
	})
}

func TestAccDataSourceLookup_metadata(t *testing.T) {
	resultJSON, err := json.Marshal(fixtures.LookupMetadataResult.Payload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_metadata,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(resultJSON)),
				),
			},
		},
	})
}

const testAccDataSourceLookup_basic = `
  data "jerakia_lookup" "lookup_1" {
    key       = "cities"
    namespace = "test"
  }
`

const testAccDataSourceLookup_singleBool = `
  data "jerakia_lookup" "lookup_1" {
    key       = "booltrue"
    namespace = "test"
  }
`

const testAccDataSourceLookup_metadata = `
  data "jerakia_lookup" "lookup_1" {
    key       = "users"
    namespace = "test"

    metadata {
      hostname = "example",
    }
  }
`
