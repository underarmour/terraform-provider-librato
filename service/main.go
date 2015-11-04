package service

import "github.com/hashicorp/terraform/helper/schema"

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
		Create: doCreate,
		Read:   doRead,
		Update: doUpdate,
		Delete: doDelete,
		Exists: doExists,
	}
}
