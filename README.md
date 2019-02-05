# terraform-provider-jerakia

Jerakia provider for Terraform

## Prerequisites

* [Terraform][1]

## Terraform Configuration Example

```hcl
provider "jerakia" {
  api_url   = "http://127.0.0.1:9843"
  api_token = "tokentoken"
}

data "jerakia_lookup" "lookup_1" {
  key       = "cities"
  namespace = "default"
}
```

## Installation

### Using a Pre-Built Binary

Downloading and installing a pre-compiled `terraform-provider-jerakia` release
is the recommended method of installation since it requires no additional tools
or libraries to be installed on your workstation.

1. Visit the [releases][2] page and download the latest release for your target
   architecture.

2. Unzip the downloaded file and copy the `terraform-provider-jerakia` binary
   to a designated directory as described in Terraform's [plugin installation
   instructions][3].

### Building from Source

> Note: Terraform requires Go 1.11 or later to successfully compile.

1. Follow these [instructions][4] to setup a Golang development environment.
2. Run:

```shell
$ go get -v -u github.com/jerakia/terraform-provider-jerakia
$ cd $GOPATH/src/github.com/jerakia/terraform-provider-jerakia
$ make build
```

You should now have a `terraform-provider-jerakia` binary located at
`$GOPATH/bin/terraform-provider-jerakia`. Copy this binary to a designated
directory as described in Terraform's [plugin installation instructions][3]

## Development

This project is using [Go Modules][5] for vendor support.

## Documentation

Full documentation can be found in the [`docs`][6] directory.

[1]: http://terraform.io
[2]: https://github.com/jerakia/terraform-provider-jerakia/releases
[3]: https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin
[4]: https://golang.org/doc/install
[5]: https://github.com/golang/go/wiki/Modules
[6]: /docs
