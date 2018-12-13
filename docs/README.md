# terraform-provider-jerakia

Use Terraform to query Jerakia

### Data Sources

* A list of supported data sources can be found [here](data_sources).

## Basic Example

To configure this provider, do the following:

```hcl
provider "jerakia" {
  api_url   = "http://127.0.0.1:9843"
  api_token = "tokentoken"
}
```

## Configuration Reference

The following arguments are supported:

* `api_url` - *Required* - The URL to the Jerakia service. This can also be set
  with the `JERAKIA_URL` environment variable.

* `api_token` - *Required* - The token to authenticate with. This can also be
	set with the `JERAKIA_TOKEN` environment variable.
