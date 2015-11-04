package request

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
)

func DoExists(d *schema.ResourceData, ip interface{}, resourceName, path string) (bool, error) {
	log.Printf("[DEBUG] doExists %s", resourceName)

	p := ip.(*provider.Provider)

	statusCode, err := DoRequest(
		"GET",
		path,
		p,
		nil,
		nil,
		200,
	)
	if err != nil {
		if statusCode == 404 {
			log.Printf("[DEBUG] doExists %s not found %s", resourceName, path)
			return false, nil
		} else {
			return false, fmt.Errorf("doExists %s failed: %v", resourceName, err)
		}
	}

	log.Printf("[DEBUG] doExists %s exists", resourceName)
	return true, nil
}
