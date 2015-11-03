package space_chart

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doExists(d *schema.ResourceData, ip interface{}) (bool, error) {
	log.Printf("[DEBUG] doExists space_chart")

	p := ip.(*provider.Provider)

	statusCode, err := request.DoRequest(
		"GET",
		fmt.Sprintf("/spaces/%s/charts/%s", d.Get("space").(string), d.Id()),
		p,
		nil,
		nil,
		200,
	)
	if err != nil {
		if statusCode == 404 {
			log.Printf("[DEBUG] doExists space_chart not found")
			return false, nil
		} else {
			return false, fmt.Errorf("doExists space_chart failed: %v", err)
		}
	}

	log.Printf("[DEBUG] doExists space_chart exists")
	return true, nil
}
