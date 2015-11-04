package service

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doUpdate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doUpdate service")

	p := ip.(*provider.Provider)
	body := &service{
		Type:     d.Get("type").(string),
		Title:    d.Get("title").(string),
		Settings: d.Get("settings").(map[string]interface{}),
	}

	_, err := request.DoRequest(
		"PUT",
		fmt.Sprintf("/services/%s", d.Id()),
		p,
		body,
		nil,
		204,
	)
	if err != nil {
		return fmt.Errorf("doUpdate service failed: %v", err)
	}

	log.Printf("[DEBUG] doUpdate service")
	return nil
}
