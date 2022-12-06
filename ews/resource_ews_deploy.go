package ews

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
	"time"
)

func resourceEwsDeploy() *schema.Resource {
	return &schema.Resource{
		Create: resourceEwsDeployUpdate,
		Read:   resourceEwsDeployRead,
		Update: resourceEwsDeployUpdate,
		Delete: resourceEwsDeployDelete,
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
			"lambda_name": {
				Description: "lambda name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"filter_path": {
				Description: "lambda zip",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "/",
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

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	log.Printf("[INFO] Created EWS with ID: %s\n", d.Id())

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
