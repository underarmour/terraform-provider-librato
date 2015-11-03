package space

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doExists(d *schema.ResourceData, ip interface{}) (bool, error) {
	log.Printf("[DEBUG] doExists space")

	p := ip.(*provider.Provider)
	s := readSpace(d)
	resp := &readResponse{}

	sc, err := request.DoRequest("GET", fmt.Sprintf("/spaces/%s", s.id), p, nil, resp, 200)
	if err != nil {
		if sc == 404 {
			log.Printf("[DEBUG] doExists space not found")
			return false, nil
		} else {
			return false, fmt.Errorf("doRead failed: %v", err)
		}
	}

	log.Printf("[DEBUG] doExists found space: %#v", resp)
	return true, nil
}
