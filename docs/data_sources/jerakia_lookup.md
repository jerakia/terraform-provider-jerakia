# jerakia_lookup

Query Jerakia.

## Basic Example

```hcl
data "jerakia_lookup" "lookup_1" {
  key       = "cities"
  namespace = "default"
}
```

## Argument Reference

* `key` - *Required* - The name of the Jerakia key to look up.

* `namespace` - *Optional* - The namespace to query. Defaults to `default`.

* `policy` - *Optional* - The policy to use for the query.

* `lookup_type` - *Optional* - The type of lookup. Valid values are `first`
	and `cascade`.

* `merge` - *Optional* - The merge strategy to use. Valid values are `array`,
	`deep_hash`, and `hash`.

* `scope` - *Optional* - Provide an alternative scope for the query.

* `scope_options` - *Optional* - A set of key/value pairs to set as parameters
	of the scope.

* `metadata` - *Optional* - A set of key/value pairs of metadata.


## Attribute Reference

* `status` - The status of the query. An error will be returned if this
	is `failed`.

* `found` - A boolean true/false if a result was found.

* `result_json` - The result of the query. The format is a JSON structure
	encoded as a string.
