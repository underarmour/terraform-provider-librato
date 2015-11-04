package space_chart

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func NewResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"space": &schema.Schema{
				Type:        schema.TypeString,
				Description: "ID of the space to add this chart to",
				Required:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Title of the chart when it is displayed",
				Required:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Indicates the type of chart. Must be one of line or stacked (default to line)",
				Optional:    true,
				Default:     "line",
			},
			"min": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The minimum display value of the chart's Y-axis",
				Optional:    true,
			},
			"max": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum display value of the chart's Y-axis",
				Optional:    true,
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The Y-axis label",
				Optional:    true,
			},
			"related_space": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of another space to which this chart is related",
				Optional:    true,
			},
			"stream": &schema.Schema{
				Type:        schema.TypeList,
				Description: "An array of hashes describing the metrics and sources to use for data in the chart.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Each stream has a unique numeric ID",
							Computed:    true,
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "A display name to use for the stream when generating the tooltip",
							Required:    true,
						},
						"metric": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Name of metric",
							Optional:    true,
						},
						"source": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Name of source or * to include all sources. This field will also accept specific wildcard entries",
							Optional:    true,
						},
						"group_function": &schema.Schema{
							Type:        schema.TypeString,
							Description: "How to process the results when multiple sources will be returned. Value must be one of average, sum, breakout",
							Optional:    true,
						},
						"composite": &schema.Schema{
							Type:        schema.TypeString,
							Description: "A composite metric query string to execute when this stream is displayed. This can not be specified with a metric, source or group_function",
							Optional:    true,
						},
						"summary_function": &schema.Schema{
							Type:        schema.TypeString,
							Description: "When visualizing complex measurements or a rolled-up measurement, this allows you to choose which statistic to use. If unset, defaults to average. Valid options are one of: [max, min, average, sum, count]",
							Optional:    true,
						},
						"color": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Sets a color to use when rendering the stream. Must be a seven character string that represents the hex code of the color e.g. #52D74C",
							Optional:    true,
						},
						"units_short": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Unit value string to use as the tooltip label",
							Optional:    true,
						},
						"units_long": &schema.Schema{
							Type:        schema.TypeString,
							Description: "String value to set as they Y-axis label. All streams that share the same units_long value will be plotted on the same Y-axis",
							Optional:    true,
						},
						"min": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Theoretical minimum Y-axis value",
							Optional:    true,
						},
						"max": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Theoretical maximum Y-axis value",
							Optional:    true,
						},
						"transform_function": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Linear formula to run on each measurement prior to visualizaton",
							Optional:    true,
						},
						"period": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "An integer value of seconds that defines the period this stream reports at",
							Optional:    true,
						},
					},
				},
			},
		},
		Create: request.CreatorFunc("space_chart", "/spaces/%s/charts", createPathFormatter, makeBody),
		Read:   request.ReaderFunc("space_chart", "/spaces/%s/charts/%s", pathFormatter, readBody),
		Update: request.UpdaterFunc("space_chart", "/spaces/%s/charts/%s", pathFormatter, makeBody),
		Delete: request.DeleterFunc("space_chart", "/spaces/%s/charts/%s", pathFormatter),
		Exists: request.ExisterFunc("space_chart", "/spaces/%s/charts/%s", pathFormatter),
	}
}

func createPathFormatter(path string, d *schema.ResourceData) string {
	return fmt.Sprintf(path, d.Get("space"))
}

func pathFormatter(path string, d *schema.ResourceData) string {
	return fmt.Sprintf(path, d.Get("space"), d.Id())
}

func makeBody(d *schema.ResourceData) map[string]interface{} {
	body := make(map[string]interface{})
	body["name"] = d.Get("name")
	body["type"] = d.Get("type")
	body["label"] = d.Get("label")
	body["related_space"] = d.Get("related_space")

	// skip empty values

	min := d.Get("min").(int)
	if min != 0 {
		body["min"] = min
	}

	max := d.Get("max").(int)
	if max != 0 {
		body["max"] = max
	}

	// clean stream empty values

	streams := d.Get("stream").([]interface{})
	for _, streamI := range streams {
		stream := streamI.(map[string]interface{})

		delete(stream, "id")

		metric := stream["metric"].(string)
		if metric == "" {
			delete(stream, "metric")
		}

		source := stream["source"].(string)
		if source == "" {
			delete(stream, "source")
		}

		group_function := stream["group_function"].(string)
		if group_function == "" {
			delete(stream, "group_function")
		}

		composite := stream["composite"].(string)
		if composite == "" {
			delete(stream, "composite")
		}

		summary_function := stream["summary_function"].(string)
		if summary_function == "" {
			delete(stream, "summary_function")
		}

		color := stream["color"].(string)
		if color == "" {
			delete(stream, "color")
		}

		units_short := stream["units_short"].(string)
		if units_short == "" {
			delete(stream, "units_short")
		}

		units_long := stream["units_long"].(string)
		if units_long == "" {
			delete(stream, "units_long")
		}

		min := stream["min"].(int)
		if min == 0 {
			delete(stream, "min")
		}

		max := stream["max"].(int)
		if max == 0 {
			delete(stream, "max")
		}

		transform_function := stream["transform_function"].(string)
		if transform_function == "" {
			delete(stream, "transform_function")
		}

		period := stream["period"].(int)
		if period == 0 {
			delete(stream, "period")
		}
	}
	body["streams"] = streams

	return body
}

func readBody(d *schema.ResourceData, resp map[string]interface{}) {
	d.Set("name", resp["name"])
	d.Set("type", resp["type"])
	d.Set("min", resp["min"])
	d.Set("max", resp["max"])
	d.Set("label", resp["label"])
	d.Set("related_space", resp["related_space"])
	d.Set("stream", resp["streams"])
}
