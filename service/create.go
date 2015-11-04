package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doCreate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doCreate new service")

	p := ip.(*provider.Provider)
	body := &service{
		Type:     d.Get("type").(string),
		Title:    d.Get("title").(string),
		Settings: d.Get("settings").(map[string]interface{}),
	}

	_, err := request.DoRequest(
		"POST",
		"/services",
		p,
		body,
		body,
		201,
	)
	if err != nil {
		return fmt.Errorf("doCreate service failed: %v", err)
	}

	log.Printf("[DEBUG] doCreate service: %#v", body)
	d.SetId(strconv.Itoa(body.Id))
	return nil
}
