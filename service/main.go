package service

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

type service struct {
	Id       int                    `json:"id,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Title    string                 `json:"title,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

func NewResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The service type (e.g. campfire, pagerduty, mail, etc.)",
				Required:    true,
			},
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Display title for the service",
				Required:    true,
			},
			"settings": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "Hash of settings specific to the service type",
				Required:    true,
			},
		},
		Create: request.CreatorFunc("service", "/services", nil, makeBody),
		Read:   request.ReaderFunc("service", "/services/%s", request.IdPathFormatter, readBody),
		Update: request.UpdaterFunc("service", "/services/%s", request.IdPathFormatter, makeBody),
		Delete: request.DeleterFunc("service", "/services/%s", request.IdPathFormatter),
		Exists: request.ExisterFunc("service", "/services/%s", request.IdPathFormatter),
	}
}

func makeBody(d *schema.ResourceData) map[string]interface{} {
	body := make(map[string]interface{})
	body["type"] = d.Get("type")
	body["title"] = d.Get("title")
	body["settings"] = d.Get("settings")
	return body
}

func readBody(d *schema.ResourceData, resp map[string]interface{}) {
	d.Set("type", resp["type"])
	d.Set("title", resp["title"])
	d.Set("settings", resp["settings"])
}
