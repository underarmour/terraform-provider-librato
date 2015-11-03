package space

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

type readResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func doRead(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doRead space")

	p := ip.(*provider.Provider)
	s := readSpace(d)
	resp := &readResponse{}

	_, err := request.DoRequest("GET", fmt.Sprintf("/spaces/%s", s.id), p, nil, resp, 200)
	if err != nil {
		return fmt.Errorf("doRead failed: %v", err)
	}

	log.Printf("[DEBUG] doRead read space: %#v", resp)
	d.SetId(strconv.Itoa(resp.Id))
	d.Set("name", resp.Name)
	return nil
}
