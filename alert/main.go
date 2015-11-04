package alert

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func NewResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A unique name used to identify the alert",
				Required:    true,
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Identifies the alert as v1 or v2",
				Optional:    true,
				Default:     2,
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A string describing this alert",
				Optional:    true,
			},
			"active": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Boolean: identifies whether the alert is active (can be triggered). Defaults to true",
				Optional:    true,
				Default:     true,
			},
			"rearm_seconds": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Specifies the minimum amount of time between sending alert notifications, in seconds. Required to be a multiple of 60, and when unset or null will default to 600",
				Optional:    true,
				Default:     600,
			},
			"conditions": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Note an alert will fire when ALL alert conditions are met",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "One of above, absent, or below",
							Required:    true,
						},
						"metric_name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The name of the metric this alert condition applies to",
							Required:    true,
						},
						"source": &schema.Schema{
							Type:        schema.TypeString,
							Description: "A source expression which identifies which sources for the given metric to monitor. If not specified all sources will be monitored",
							Optional:    true,
						},
						"threshold": &schema.Schema{
							Type:        schema.TypeFloat,
							Description: "Measurements over this number will fire the alert.",
							Optional:    true,
						},
						"summary_function": &schema.Schema{
							Type:        schema.TypeString,
							Description: "For gauge metrics will default to average, which is also the value of non-complex or un-aggregated measurements. If set, must be one of: [min, max, average, sum, count, derivative]",
							Optional:    true,
						},
						"duration": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Number of seconds that data for the specified metric/source combination must be above the threshold for before the condition is met. If unset, a single sample above the threshold will trigger the condition",
							Optional:    true,
						},
						"detect_reset": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "If the summary_function is derivative, this toggles the method used to calculate the delta from the previous sample. When set to false (default), the delta is calculated as simple subtraction of current - previous. If true only increasing (positive) values will be reported",
							Optional:    true,
							Default:     false,
						},
					},
				},
			},
			"services": &schema.Schema{
				Type:        schema.TypeList,
				Description: "An array of services to notify for this alert",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"attributes": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "A key-value hash of metadata for the alert",
				Optional:    true,
			},
		},
		Create: request.CreatorFunc("alert", "/alerts", nil, makeBody),
		Read:   request.ReaderFunc("alert", "/alerts/%s", request.IdPathFormatter, readBody),
		Update: request.UpdaterFunc("alert", "/alerts/%s", request.IdPathFormatter, makeBody),
		Delete: request.DeleterFunc("alert", "/alerts/%s", request.IdPathFormatter),
		Exists: request.ExisterFunc("alert", "/alerts/%s", request.IdPathFormatter),
	}
}

func makeBody(d *schema.ResourceData) map[string]interface{} {
	body := make(map[string]interface{})

	body["name"] = d.Get("name")
	body["version"] = d.Get("version")
	body["active"] = d.Get("active")
	body["rearm_seconds"] = d.Get("rearm_seconds")
	body["services"] = d.Get("services")
	body["description"] = d.Get("description").(string)

	// default runbook_url

	attributes := d.Get("attributes").(map[string]interface{})
	if _, ok := attributes["runbook_url"]; !ok {
		attributes["runbook_url"] = ""
	}
	body["attributes"] = attributes

	// clean empty conditions values

	conditions := d.Get("conditions").([]interface{})
	for _, conditionI := range conditions {
		condition := conditionI.(map[string]interface{})

		source := condition["source"].(string)
		if source == "" {
			delete(condition, "source")
		}

		summaryFunction := condition["summary_function"].(string)
		if summaryFunction == "" {
			delete(condition, "summary_function")
		}

		duration := condition["duration"].(int)
		if duration == 0 {
			delete(condition, "duration")
		}
	}
	body["conditions"] = conditions

	return body
}

func readBody(d *schema.ResourceData, resp map[string]interface{}) {
	d.Set("name", resp["name"])
	d.Set("version", resp["version"])
	d.Set("description", resp["description"])
	d.Set("active", resp["active"])
	d.Set("rearm_seconds", resp["rearm_seconds"])
	d.Set("conditions", resp["conditions"])
	d.Set("attributes", resp["attributes"])

	// pull out services ids because the request body
	// is different than the response body

	services := resp["services"].([]interface{})
	serviceIds := make([]int, 0)
	for _, serviceI := range services {
		service := serviceI.(map[string]interface{})

		// not sure why this is introspecting as a float64
		serviceIds = append(serviceIds, int(service["id"].(float64)))
	}
	d.Set("services", serviceIds)
}
