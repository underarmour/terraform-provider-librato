package request

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
)

func doRead(d *schema.ResourceData, ip interface{}, resourceName, path string, readBody readBodyFn) error {
	log.Printf("[DEBUG] doRead %s", resourceName)

	p := ip.(*provider.Provider)
	var resp map[string]interface{}

	_, err := DoRequest(
		"GET",
		path,
		p,
		nil,
		&resp,
		200,
	)

	if err != nil {
		return fmt.Errorf("doRead %s failed: %v", resourceName, err)
	}

	log.Printf("[DEBUG] doRead %s: %#v", resourceName, resp)
	readBody(d, resp)
	return nil
}

func ReaderFunc(
	resourceName string,
	path string,
	pathFormatter pathFormatterFn,
	readBody readBodyFn,
) schema.ReadFunc {
	return func(d *schema.ResourceData, ip interface{}) error {
		if pathFormatter != nil {
			path = pathFormatter(path, d)
		}

		return doRead(d, ip, resourceName, path, readBody)
	}
}
