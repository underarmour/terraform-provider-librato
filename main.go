package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/space"
	"github.com/underarmour/terraform-provider-librato/space_chart"
)

func configureFunc(d *schema.ResourceData) (interface{}, error) {
	user := d.Get("user").(string)
	token := d.Get("token").(string)
	baseUrl := fmt.Sprintf("https://%v:%v@metrics-api.librato.com/v1", user, token)
	p := provider.NewProvider(user, token, baseUrl)
	log.Printf("[DEBUG] Configured new provider struct: %#v", p)
	return p, nil
}

func newProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "User is the email address that you used to create your Librato Metrics account",
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Token is the API token that can be found on your account page",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"librato_space":       space.NewResource(),
			"librato_space_chart": space_chart.NewResource(),
		},
		ConfigureFunc: configureFunc,
	}
}

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: newProvider})
}
