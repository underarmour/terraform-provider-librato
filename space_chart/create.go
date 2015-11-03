package space_chart

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

type stream struct {
	Id                int    `json:"id,omitempty"`
	Name              string `json:"name"`
	Metric            string `json:"metric,omitempty"`
	Source            string `json:"source,omitempty"`
	GroupFunction     string `json:"group_function,omitempty"`
	Composite         string `json:"composite,omitempty"`
	SummaryFunction   string `json:"summary_function,omitempty"`
	Color             string `json:"color,omitempty"`
	UnitsShort        string `json:"units_short,omitempty"`
	UnitsLong         string `json:"units_long,omitempty"`
	Min               int    `json:"min,omitempty"`
	Max               int    `json:"max,omitempty"`
	TransformFunction string `json:"transform_function,omitempty"`
	Period            int    `json:"period,omitempty"`
}

type createBody struct {
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Min          int       `json:"min,omitempty"`
	Max          int       `json:"max,omitempty"`
	Label        string    `json:"label,omitempty"`
	RelatedSpace string    `json:"related_space,omitempty"`
	Streams      []*stream `json:"streams,omitempty"`
}

type createResponse struct {
	Id int `json:"id"`
}

func makeStreams(streamsI []interface{}) []*stream {
	streams := []*stream{}
	for _, streamI := range streamsI {
		streamM := streamI.(map[string]interface{})
		streams = append(streams, &stream{
			Name:              streamM["name"].(string),
			Metric:            streamM["metric"].(string),
			Source:            streamM["source"].(string),
			GroupFunction:     streamM["group_function"].(string),
			Composite:         streamM["composite"].(string),
			SummaryFunction:   streamM["summary_function"].(string),
			Color:             streamM["color"].(string),
			UnitsShort:        streamM["units_short"].(string),
			UnitsLong:         streamM["units_long"].(string),
			Min:               streamM["min"].(int),
			Max:               streamM["max"].(int),
			TransformFunction: streamM["transform_function"].(string),
			Period:            streamM["period"].(int),
		})
	}
	return streams
}

func doCreate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doCreate new space_chart")

	p := ip.(*provider.Provider)
	body := &createBody{
		Name:         d.Get("name").(string),
		Type:         d.Get("type").(string),
		Min:          d.Get("min").(int),
		Max:          d.Get("max").(int),
		Label:        d.Get("label").(string),
		RelatedSpace: d.Get("related_space").(string),
		Streams:      makeStreams(d.Get("stream").([]interface{})),
	}

	resp := &createResponse{}
	_, err := request.DoRequest(
		"POST",
		fmt.Sprintf("/spaces/%s/charts", d.Get("space").(string)),
		p,
		body,
		resp,
		201,
	)
	if err != nil {
		return fmt.Errorf("doCreate space_chart failed: %v", err)
	}

	log.Printf("[DEBUG] doCreate space_chart: %#v", resp)
	d.SetId(strconv.Itoa(resp.Id))
	return nil
}
