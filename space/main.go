package space

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func NewResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Unique name for space",
				Required:    true,
			},
		},
		Create: request.CreatorFunc("space", "/spaces", nil, makeBody),
		Read:   request.ReaderFunc("space", "/spaces/%s", request.IdPathFormatter, readBody),
		Update: request.UpdaterFunc("space", "/spaces/%s", request.IdPathFormatter, makeBody),
		Delete: request.DeleterFunc("space", "/spaces/%s", request.IdPathFormatter),
		Exists: request.ExisterFunc("space", "/spaces/%s", request.IdPathFormatter),
	}
}

func makeBody(d *schema.ResourceData) map[string]interface{} {
	body := make(map[string]interface{})
	body["name"] = d.Get("name")
	return body
}

func readBody(d *schema.ResourceData, resp map[string]interface{}) {
	d.Set("name", resp["name"])
}
