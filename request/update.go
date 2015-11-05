package request

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
)

func doUpdate(
	d *schema.ResourceData,
	ip interface{},
	resourceName,
	path string,
	makeBody makeBodyFn,
) error {
	log.Printf("[DEBUG] doUpdate %s", resourceName)

	p := ip.(*provider.Provider)
	body := makeBody(d)

	_, err := DoRequest(
		"PUT",
		path,
		p,
		body,
		nil,
		204,
	)

	if err != nil {
		return fmt.Errorf("doUpdate %s failed: %v", resourceName, err)
	}

	log.Printf("[DEBUG] doUpdate %s", resourceName)
	return nil
}

func UpdaterFunc(
	resourceName string,
	path string,
	pathFormatter pathFormatterFn,
	makeBody makeBodyFn,
) schema.UpdateFunc {
	return func(d *schema.ResourceData, ip interface{}) error {
		return doUpdate(d, ip, resourceName, formatPath(path, pathFormatter, d), makeBody)
	}
}
