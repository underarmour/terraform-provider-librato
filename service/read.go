package service

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doRead(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doRead service")

	p := ip.(*provider.Provider)
	resp := &service{}

	_, err := request.DoRequest(
		"GET",
		fmt.Sprintf("/services/%s", d.Id()),
		p,
		nil,
		resp,
		200,
	)
	if err != nil {
		return fmt.Errorf("doRead service failed: %v", err)
	}

	log.Printf("[DEBUG] doRead service: %#v", resp)
	d.Set("type", resp.Type)
	d.Set("title", resp.Title)
	d.Set("settings", resp.Settings)
	return nil
}
