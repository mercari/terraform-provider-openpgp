package openpgp

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyCreate,
		Read:   resourceKeyRead,
		Delete: resourceKeyDelete,

		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of PGP key`,
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Email of PGP key`,
			},
			// Computed
			"private_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKeyCreate(d *schema.ResourceData, m interface{}) error {
	key, err := openpgp.NewEntity(
		d.Get("name").(string),  // name
		"",                      // comment
		d.Get("email").(string), // email
		nil,                     // use sensible defaults intentionally
	)
	if err != nil {
		return fmt.Errorf("failed to create new PGP key pair: %w", err)
	}

	armorPrivateKey, err := armorEncodeKey(key, openpgp.PrivateKeyType)
	if err != nil {
		return err
	}
	if err := d.Set("private_key", armorPrivateKey); err != nil {
		return err
	}

	armorPublicKey, err := armorEncodeKey(key, openpgp.PublicKeyType)
	if err != nil {
		return err
	}
	if err := d.Set("public_key", armorPublicKey); err != nil {
		return err
	}

	fingerprint := key.PrimaryKey.KeyIdString()
	if err := d.Set("fingerprint", fingerprint); err != nil {
		return err
	}
	d.SetId(fingerprint)

	return resourceKeyRead(d, m)
}

func resourceKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceKeyDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func armorEncodeKey(key *openpgp.Entity, keyType string) (string, error) {
	var b bytes.Buffer
	w, err := armor.Encode(&b, keyType, make(map[string]string))
	if err != nil {
		return "", fmt.Errorf("failed to create armor encoder: %w", err)
	}

	switch keyType {
	case openpgp.PrivateKeyType:
		if err := key.SerializePrivate(w, nil); err != nil {
			return "", fmt.Errorf("failed to serialize private key: %w", err)
		}
	case openpgp.PublicKeyType:
		if err := key.Serialize(w); err != nil {
			return "", fmt.Errorf("failed to serialize public key: %w", err)
		}
	default:
		return "", fmt.Errorf("unknown key type: %s", keyType)
	}
	if err := w.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer for armor: %w", err)
	}
	return b.String(), nil
}
