package alert

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

type createResponse struct {
	Id int `json:"id"`
}

func doCreate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doCreate new alert")

	p := ip.(*provider.Provider)
	body := makeBody(d)

	resp := &createResponse{}
	_, err := request.DoRequest(
		"POST",
		"/alerts",
		p,
		body,
		resp,
		201,
	)
	if err != nil {
		return fmt.Errorf("doCreate alert failed: %v", err)
	}

	log.Printf("[DEBUG] doCreate alert: %#v", resp)
	d.SetId(strconv.Itoa(resp.Id))
	return nil
}
