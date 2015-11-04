package space

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doExists(d *schema.ResourceData, ip interface{}) (bool, error) {
	return request.DoExists(
		d, ip, "space",
		fmt.Sprintf("/spaces/%s", d.Id()),
	)
}
