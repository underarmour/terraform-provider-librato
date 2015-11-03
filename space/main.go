package space

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

type space struct {
	id   string
	name string
}

func readSpace(d *schema.ResourceData) *space {
	id := d.Get("id").(string)
	name := d.Get("name").(string)
	s := &space{
		id:   id,
		name: name,
	}
	log.Printf("[DEBUG] readSpace from resource data: %#v", s)
	return s
}

func doUpdate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] one")
	return nil
}

func doDelete(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] one")
	return nil
}

func NewResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Each space has a unique numeric ID",
				Computed:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Unique name for space",
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
