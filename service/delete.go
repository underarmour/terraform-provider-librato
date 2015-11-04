package service

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doDelete(d *schema.ResourceData, ip interface{}) error {
	return request.DoDelete(
		d, ip, "service",
		fmt.Sprintf("/services/%s", d.Id()),
	)
}
