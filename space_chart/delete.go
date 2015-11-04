package space_chart

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doDelete(d *schema.ResourceData, ip interface{}) error {
	return request.DoDelete(
		d, ip, "space_chart",
		fmt.Sprintf("/spaces/%s/charts/%s", d.Get("space").(string), d.Id()),
	)
}
