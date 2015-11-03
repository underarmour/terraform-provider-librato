package space

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doDelete(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doDelete space")

	p := ip.(*provider.Provider)

	_, err := request.DoRequest(
		"DELETE",
		fmt.Sprintf("/spaces/%s", d.Id()),
		p,
		nil,
		nil,
		204,
	)
	if err != nil {
		return fmt.Errorf("doDelete space failed: %v", err)
	}

	log.Printf("[DEBUG] doDelete deleted space")
	return nil
}