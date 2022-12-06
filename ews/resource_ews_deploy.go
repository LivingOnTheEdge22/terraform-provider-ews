package ews

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceEwsDeploy() *schema.Resource {
	return &schema.Resource{
		Create: resourceEwsDeployUpdate,
		Read:   resourceEwsDeployRead,
		Update: resourceEwsDeployUpdate,
		Delete: resourceEwsDeployDelete,

		Schema: map[string]*schema.Schema{
			// Required Arguments
			"account_id": {
				Description: "account id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"site_id": {
				Description: "Site id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"lambda_name": {
				Description: "lambda name",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"filter_path": {
				Description: "lambda zip",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "/",
				ForceNew:    true,
			},
		},
	}
}

func resourceEwsDeployUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeployWebAssembly(
		d.Get("account_id").(string),
		d.Get("lambda_name").(string),
		d.Get("filter_path").(string),
	)
	if err != nil {
		return err
	}

	syntheticId := fmt.Sprintf("%s-%s", d.Get("lambda_name").(string), d.Get("site_id").(string))
	d.SetId(syntheticId)

	log.Printf("[INFO] Deployed EWS with ID: %s\n", d.Id())

	return resourceEwsDeployRead(d, m)
}

func resourceEwsDeployRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEwsDeployDelete(d *schema.ResourceData, m interface{}) error {
	// Set the ID to empty
	// Implicitly clears the resource
	d.SetId("")
	return nil
}
