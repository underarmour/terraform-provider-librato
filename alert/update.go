package alert

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doUpdate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doUpdate alert")

	p := ip.(*provider.Provider)
	body := makeBody(d)

	_, err := request.DoRequest(
		"PUT",
		fmt.Sprintf("/alerts/%s", d.Id()),
		p,
		body,
		nil,
		204,
	)
	if err != nil {
		return fmt.Errorf("doUpdate alert failed: %v", err)
	}

	log.Printf("[DEBUG] doUpdate alert")
	return nil
}
