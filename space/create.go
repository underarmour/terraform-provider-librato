package space

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

type createBody struct {
	Name string `json:"name"`
}

func doCreate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doCreate new space")

	p := ip.(*provider.Provider)
	s := readSpace(d)
	body := &createBody{Name: s.name}

	resp := &readResponse{}
	_, err := request.DoRequest("POST", "/spaces", p, body, resp, 201)
	if err != nil {
		return fmt.Errorf("doCreate space failed: %v", err)
	}

	log.Printf("[DEBUG] doCreate created space: %#v", resp)
	d.SetId(strconv.Itoa(resp.Id))
	d.Set("name", resp.Name)
	return nil
}
