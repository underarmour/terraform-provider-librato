package alert

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doRead(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doRead alert")

	p := ip.(*provider.Provider)
	var resp map[string]interface{}

	_, err := request.DoRequest(
		"GET",
		fmt.Sprintf("/alerts/%s", d.Id()),
		p,
		nil,
		&resp,
		200,
	)
	if err != nil {
		return fmt.Errorf("doRead alert failed: %v", err)
	}

	log.Printf("[DEBUG] doRead alert: %#v", resp)
	d.Set("name", resp["name"])
	d.Set("version", resp["version"])
	d.Set("description", resp["description"])
	d.Set("active", resp["active"])
	d.Set("rearm_seconds", resp["rearm_seconds"])
	d.Set("conditions", resp["conditions"])
	d.Set("services", resp["services"])
	d.Set("attributes", resp["attributes"])
	return nil
}
