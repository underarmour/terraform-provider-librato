package space_chart

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doUpdate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doUpdate space_chart")

	p := ip.(*provider.Provider)
	body := &createBody{
		Name:         d.Get("name").(string),
		Type:         d.Get("type").(string),
		Min:          d.Get("min").(int),
		Max:          d.Get("max").(int),
		Label:        d.Get("label").(string),
		RelatedSpace: d.Get("related_space").(string),
		Streams:      makeStreams(d.Get("stream").([]interface{})),
	}

	_, err := request.DoRequest(
		"PUT",
		fmt.Sprintf("/spaces/%s/charts/%s", d.Get("space").(string), d.Id()),
		p,
		body,
		nil,
		204,
	)
	if err != nil {
		return fmt.Errorf("doUpdate space_chart failed: %v", err)
	}

	log.Printf("[DEBUG] doUpdate space_chart")
	return nil
}
