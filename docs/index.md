---
Page title: "Provider: Psychz"
---

# Psychz Provider

The Psychz provider is used to interact with the resources provided by Psychz Platform. The provider needs to be configured with the proper credentials before
it can be used.

For information about obtaining API key and secret required for Psychz Details and Order Express refer to
[Generating API UserName and Access Token](https://www.psychz.net/dashboard/client/web/account/api)

Interacting with Psychz Order Express and Order Details requires an API access token that can be generated at Project-level or User-level tokens can be used.

## Example Usage

Example HCL with [provider configuration](https://www.terraform.io/docs/configuration/providers.html)
and a [required providers definition](https://www.terraform.io/language/settings#specifying-a-required-terraform-version):

```hcl
terraform {
  required_providers {
    psychz = {
      source = "psychz.net/psychz/psychz"
      version = "1.0.0"
    }
  }
}
