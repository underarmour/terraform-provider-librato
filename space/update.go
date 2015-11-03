package space

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doUpdate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doUpdate space")

	p := ip.(*provider.Provider)
	body := &createBody{Name: d.Get("name").(string)}

	_, err := request.DoRequest("PUT", fmt.Sprintf("/spaces/%s", d.Id()), p, body, nil, 204)
	if err != nil {
		return fmt.Errorf("doUpdate space failed: %v", err)
	}

	log.Printf("[DEBUG] doUpdate updated space")
	return nil
}
