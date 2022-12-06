package ews

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceEwsCompile() *schema.Resource {
	return &schema.Resource{
		Create: resourceEwsCompileUpdate,
		Read:   resourceEwsCompileRead,
		Update: resourceEwsCompileUpdate,
		Delete: resourceEwsCompileDelete,
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
			"lambda_path": {
				Description: "lambda zip",
				Type:        schema.TypeString,
				Required:    true,
			},
			"lambda_name": {
				Description: "lambda name",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceEwsCompileUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.CompileWebAssembly(
		d.Get("account_id").(string),
		d.Get("lambda_name").(string),
		d.Get("lambda_path").(string),
	)

	if err != nil {
		return err
	}
	//todo Raphy - new id generation
	id := "generated string"
	d.SetId(id)
	log.Printf("[INFO] Created EWS with ID: %s\n", d.Id())

	return resourceEwsCompileRead(d, m)
}

func resourceEwsCompileRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEwsCompileDelete(d *schema.ResourceData, m interface{}) error {
	// Set the ID to empty
	// Implicitly clears the resource
	d.SetId("")
	return nil
}
