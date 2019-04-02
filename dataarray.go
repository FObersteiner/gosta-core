package core

// CreateObservations is used to POST multiple observations at once
type CreateObservations struct {
	Datastream *Datastream      `json:"Datastream,omitempty"`
	Components []string         `json:"components,omitempty"`
	Count      int64            `json:"dataArray@iot.count,omitempty"`
	Data       [][]*interface{} `json:"dataArray,omitempty"`
}
