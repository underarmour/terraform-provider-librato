package request

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
)

func doDelete(
	d *schema.ResourceData,
	ip interface{},
	resourceName,
	path string,
) error {
	log.Printf("[DEBUG] doDelete %s", resourceName)

	p := ip.(*provider.Provider)

	_, err := DoRequest(
		"DELETE",
		path,
		p,
		nil,
		nil,
		204,
	)

	if err != nil {
		return fmt.Errorf("doDelete %s failed: %v", resourceName, err)
	}

	log.Printf("[DEBUG] doDelete %s", resourceName)
	return nil
}

func DeleterFunc(
	resourceName string,
	path string,
	pathFormatter pathFormatterFn,
) schema.DeleteFunc {
	return func(d *schema.ResourceData, ip interface{}) error {
		return doDelete(d, ip, resourceName, formatPath(path, pathFormatter, d))
	}
}
