# Terraform OpenPGP provider

This provider creates a OpenPGP key pair so that you can use the key pair for other Terraform resources.

## Requirements

- Terraform 0.11.x or higher
- Go 1.13 or higher (for build)

## Using the provider

Download the latest release for your OS from the [release page](https://github.com/mercari/terraform-provider-openpgp/releases) and follow the instructions to [install third party plugins](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

*Example*:
``` hcl
# Create a OpenPGP key
resource "openpgp_key" "my-openpgp-key" {
  name  = "Your name"
  email = "you@example.com"
}
```

You can use the following attributes:

- public_key - The public key of the OpenPGP key pair. (ASCII-armored format)
- private_key - The private key of the OpenPGP key pair. (ASCII-armored format)
- fingerprint - The fingerprint of the OpenPGP key pair.

## Contribution

Please read the CLA below carefully before submitting your contribution.

https://www.mercari.com/cla/

