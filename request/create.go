package request

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
)

type createResponse struct {
	Id int `json:"id"`
}

func doCreate(
	d *schema.ResourceData,
	ip interface{},
	resourceName string,
	path string,
	makeBody makeBodyFn,
) error {
	log.Printf("[DEBUG] doCreate new %s", resourceName)

	p := ip.(*provider.Provider)
	body := makeBody(d)
	resp := &createResponse{}

	_, err := DoRequest(
		"POST",
		path,
		p,
		body,
		resp,
		201,
	)

	if err != nil {
		return fmt.Errorf("doCreate %s failed: %v", resourceName, err)
	}

	log.Printf("[DEBUG] doCreate %s: %#v", resourceName, resp)
	d.SetId(strconv.Itoa(resp.Id))
	return nil
}

func CreatorFunc(
	resourceName string,
	path string,
	pathFormatter pathFormatterFn,
	makeBody makeBodyFn,
) schema.CreateFunc {
	return func(d *schema.ResourceData, ip interface{}) error {
		return doCreate(d, ip, resourceName, formatPath(path, pathFormatter, d), makeBody)
	}
}
