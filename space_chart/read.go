package space_chart

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

func makeMaps(streams []*stream) []map[string]interface{} {
	maps := []map[string]interface{}{}
	for _, stream := range streams {
		maps = append(maps, map[string]interface{}{
			"id":                 stream.Id,
			"name":               stream.Name,
			"metric":             stream.Metric,
			"source":             stream.Source,
			"group_function":     stream.GroupFunction,
			"composite":          stream.Composite,
			"summary_function":   stream.SummaryFunction,
			"color":              stream.Color,
			"units_short":        stream.UnitsShort,
			"units_long":         stream.UnitsLong,
			"min":                stream.Min,
			"max":                stream.Max,
			"transform_function": stream.TransformFunction,
			"period":             stream.Period,
		})
	}
	return maps
}

func doRead(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doRead space_chart")

	p := ip.(*provider.Provider)
	resp := &createBody{}

	_, err := request.DoRequest(
		"GET",
		fmt.Sprintf("/spaces/%s/charts/%s", d.Get("space").(string), d.Id()),
		p,
		nil,
		resp,
		200,
	)
	if err != nil {
		return fmt.Errorf("doRead space_chart failed: %v", err)
	}

	log.Printf("[DEBUG] doRead space_chart: %#v", resp)
	d.Set("name", resp.Name)
	d.Set("type", resp.Type)
	d.Set("min", resp.Min)
	d.Set("max", resp.Max)
	d.Set("label", resp.Label)
	d.Set("related_space", resp.RelatedSpace)
	d.Set("stream", makeMaps(resp.Streams))
	return nil
}
