package openpgp

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCreateKey(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "openpgp_key" "test1" {
}
`,
				ExpectError: regexp.MustCompile(`The argument "name" is required, but no definition was found.`),
			},
			{
				Config: `
resource "openpgp_key" "test2" {
  name = "Dragon3"
}
`,
				ExpectError: regexp.MustCompile(`The argument "email" is required, but no definition was found.`),
			},
			{
				Config: `
resource "openpgp_key" "test3" {
  email = "dragon3@example.com"
}
`,
				ExpectError: regexp.MustCompile(`The argument "name" is required, but no definition was found.`),
			},
			{
				Config: `
resource "openpgp_key" "test4" {
  name = "Dragon3"
  email = "dragon3@example.com"
}

output "private_key" {
  value     = "${openpgp_key.test4.private_key}"
  sensitive = true
}
output "public_key" {
  value     = "${openpgp_key.test4.public_key}"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("openpgp_key.test4", "name", "Dragon3"),
					resource.TestCheckResourceAttr("openpgp_key.test4", "email", "dragon3@example.com"),
					resource.TestCheckResourceAttrSet("openpgp_key.test4", "private_key"),
					resource.TestCheckResourceAttrSet("openpgp_key.test4", "public_key"),
					resource.TestCheckResourceAttrSet("openpgp_key.test4", "fingerprint"),
					resource.TestMatchResourceAttr("openpgp_key.test4", "private_key",
						regexp.MustCompile("^-----BEGIN PGP PRIVATE KEY BLOCK-----")),
					resource.TestMatchResourceAttr("openpgp_key.test4", "public_key",
						regexp.MustCompile("^-----BEGIN PGP PUBLIC KEY BLOCK-----")),
				),
			},
		},
	})
}
