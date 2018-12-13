package jerakia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/jerakia/go-jerakia"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceLookup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLookupRead,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},

			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"lookup_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"first", "cascade",
				}, false),
			},

			"merge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"array", "deep_hash", "hash",
				}, false),
			},

			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"scope_options": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			// Computed
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"found": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"result_json": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLookupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var lookupOpts jerakia.LookupOpts

	if v, ok := d.GetOkExists("namespace"); ok {
		namespace := v.(string)
		lookupOpts.Namespace = namespace
	}

	if v, ok := d.GetOkExists("policy"); ok {
		policy := v.(string)
		lookupOpts.Policy = policy
	}

	if v, ok := d.GetOkExists("lookup_type"); ok {
		lookupType := v.(string)
		lookupOpts.LookupType = lookupType
	}

	if v, ok := d.GetOkExists("merge"); ok {
		merge := v.(string)
		lookupOpts.Merge = merge
	}

	if v, ok := d.GetOkExists("scope"); ok {
		scope := v.(string)
		lookupOpts.Scope = scope
	}

	if v, ok := d.GetOkExists("scope_options"); ok {
		scopeOptions := expandMap(v.(map[string]interface{}))
		lookupOpts.ScopeOptions = scopeOptions
	}

	if v, ok := d.GetOkExists("metadata"); ok {
		metadata := expandMap(v.(map[string]interface{}))
		lookupOpts.Metadata = metadata
	}

	key := d.Get("key").(string)
	log.Printf("[DEBUG] jerakia_lookup lookup options for %s: %#v", key, lookupOpts)

	result, err := jerakia.Lookup(&config.client, key, &lookupOpts)
	if err != nil {
		return fmt.Errorf("Error querying Jerakia for %s: %s", key, err)
	}

	log.Printf("[DEBUG] jerakia_lookup result for %s: %#v", key, result)

	if result.Status == "failed" {
		return fmt.Errorf("Error querying Jerakia for %s: %s", key, result.Message)
	}

	d.SetId(generateId(lookupOpts))

	d.Set("status", result.Status)
	d.Set("found", result.Found)

	payload, err := json.Marshal(result.Payload)
	if err != nil {
		return fmt.Errorf("Error marshaling Jerakia result for %s: %s", key, err)
	}

	d.Set("result_json", string(payload))

	return nil
}

func generateId(opts jerakia.LookupOpts) string {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(opts)
	return strconv.Itoa(hashcode.String(buf.String()))
}

func expandMap(v map[string]interface{}) map[string]string {
	vmap := make(map[string]string)

	for k, v := range v {
		if value, ok := v.(string); ok && value != "" {
			vmap[k] = value
		}
	}

	return vmap
}
