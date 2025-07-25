package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateObservations is used to POST multiple observations at once
type CreateObservations struct {
	BaseEntity
	Datastreams []*Datastream `json:"Observations,omitempty"`
}

// RawCreateObservations is used to POST multiple observations at once
type RawCreateObservations struct {
	Datastream *Datastream         `json:"Datastream,omitempty"`
	Components []string            `json:"components,omitempty"`
	Count      int64               `json:"dataArray@iot.count,omitempty"`
	Data       [][]json.RawMessage `json:"dataArray,omitempty"`
}

// GetEntityType returns the EntityType for Datastream
func (c CreateObservations) GetEntityType() EntityType {
	return EntityTypeCreateObservations
}

// ParseEntity tries to parse the given json byte array into the current entity
func (c *CreateObservations) ParseEntity(data []byte) error {
	var rco []*RawCreateObservations
	err := json.Unmarshal(data, &rco)
	if err != nil {
		return errors.New("Unable to parse")
	}

	c.Datastreams = make([]*Datastream, 0)

	// foreach datastream description
	for i := 0; i < len(rco); i++ {
		if len(fmt.Sprintf("%v", rco[i].Datastream.ID)) == 0 {
			continue
		}

		rco[i].Datastream.Observations = make([]*Observation, 0)
		// for each observation
		for j := 0; j < len(rco[i].Data); j++ {
			obs := &Observation{}
			// for each value
			for k := 0; k < len(rco[i].Data[j]); k++ {
				field := strings.ToLower(rco[i].Components[k])
				val := rco[i].Data[j][k]
				var s string

				switch field {
				case "phenomenontime":
					json.Unmarshal(val, &s)
					obs.PhenomenonTime = s
				case "result":
					obs.Result = val
				case "resulttime":
					json.Unmarshal(val, &s)
					obs.ResultTime = &s
				case "resultquality":
					json.Unmarshal(val, &s)
					obs.ResultQuality = s
				case "validtime":
					json.Unmarshal(val, &s)
					obs.ValidTime = s
				case "parameters":
					var params map[string]interface{}
					json.Unmarshal(val, &params)
					obs.Parameters = params
				case "featureofinterest/id":
					var id interface{}
					json.Unmarshal(val, &id)
					foi := &FeatureOfInterest{}
					foi.ID = id
					obs.FeatureOfInterest = foi
				}
			}

			rco[i].Datastream.Observations = append(rco[i].Datastream.Observations, obs)
		}

		c.Datastreams = append(c.Datastreams, rco[i].Datastream)
	}

	return nil
}

// SetAllLinks sets the self link and relational links
func (c *CreateObservations) SetAllLinks(externalURL string) {}

// SetLinks sets the entity specific navigation links, empty string if linked(expanded) data is not nil
func (c *CreateObservations) SetLinks(externalURL string) {}

// SetSelfLink sets the self link for the entity
func (c *CreateObservations) SetSelfLink(externalURL string) {
	c.NavSelf = CreateEntitySelfLink(externalURL, EntityLinkCreateObservations.ToString(), nil)
}

// ContainsMandatoryParams checks if all mandatory params for a Datastream are available before posting
func (c *CreateObservations) ContainsMandatoryParams() (bool, []error) {
	// ToDo cheack all sub params
	return true, nil
}
