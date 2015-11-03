package space_chart

import "github.com/hashicorp/terraform/helper/schema"

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
		Create: doCreate,
		Read:   doRead,
		Update: doUpdate,
		Delete: doDelete,
		Exists: doExists,
	}
}
