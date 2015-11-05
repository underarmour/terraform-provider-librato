package request

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
)

func doRead(
	d *schema.ResourceData,
	ip interface{},
	resourceName,
	path string,
	readBody readBodyFn,
) error {
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
	errs := readBody(d, resp)

	if errs != nil {
		errstrs := make([]string, len(errs))
		for _, err := range errs {
			errstrs = append(errstrs, err.Error())
		}

		errstr := strings.Join(errstrs, "; ")
		return fmt.Errorf("doRead %s readBody failed: %s", resourceName, errstr)
	}

	return nil
}

func ReaderFunc(
	resourceName string,
	path string,
	pathFormatter pathFormatterFn,
	readBody readBodyFn,
) schema.ReadFunc {
	return func(d *schema.ResourceData, ip interface{}) error {
		return doRead(d, ip, resourceName, formatPath(path, pathFormatter, d), readBody)
	}
}
