package ews

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceEws() *schema.Resource {
	return &schema.Resource{
		Create: resourceEwsCreate,
		Read:   resourceEwsRead,
		Update: resourceEwsUpdate,
		Delete: resourceEwsDelete,
		//Importer: &schema.ResourceImporter{
		//	StateContext: schema.ImportStatePassthroughContext,
		//},
		Schema: map[string]*schema.Schema{
			// Required Arguments
			"account_id": {
				Description: "account id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"lambda": {
				Description: "lambda zip",
				Type:        schema.TypeString,
				Required:    true,
			},
			"lambda_name": {
				Description: "lambda zip",
				Type:        schema.TypeString,
				Required:    true,
			},
			"filter_path": {
				Description: "lambda zip",
				Type:        schema.TypeString,
				Required:    true,
			},
			"deployed": {
				Description: "is lambda deployed",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceEwsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	wasm := WASMStruct{
		Lambda: d.Get("lambda").(string),
	}

	ewsApiDTO := EwsApiDTO{
		Data: wasm,
	}

	_, err := client.CompileWebAssembly(
		d.Get("account_id").(string),
		ewsApiDTO,
	)

	if err != nil {
		return err
	}
	id := "generated string"
	d.SetId(id)
	log.Printf("[INFO] Created EWS with ID: %s\n", d.Id())

	deployed := d.Get("deployed").(bool)
	if d.HasChange("deployed") && d.Get("deployed") != "" {
		if deployed == true {
			_, err = client.DeployWebAssembly(
				d.Get("account_id").(string),
				d.Get("lambda_name").(string),
				d.Get("filter_path").(string),
			)
			if err != nil {
				return err
			}

		} else {
			//todo - ask Siva
		}
	}

	return resourceEwsRead(d, m)
}

func resourceEwsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEwsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceEwsRead(d, m)
}

func resourceEwsDelete(d *schema.ResourceData, m interface{}) error {
	// Set the ID to empty
	// Implicitly clears the resource
	d.SetId("")
	return nil
}
