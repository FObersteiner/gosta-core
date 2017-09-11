package core

// ArrayResponse is the default response format for sending content back used by the server
type ArrayResponse struct {
	Count    int          `json:"@iot.count,omitempty"`
	NextLink string       `json:"@iot.nextLink,omitempty"`
	Data     *interface{} `json:"value"`
}

// ArrayResponseEndpoint can be used to parse an array response containing endpoint information
type ArrayResponseEndpoint struct {
	ArrayResponse
	Data []*Endpoint `json:"value"`
}

// ArrayResponseThings can be used to parse an array response containing things
type ArrayResponseThings struct {
	ArrayResponse
	Data []*Thing `json:"value"`
}

// ArrayResponseLocations can be used to parse an array response containing locations
type ArrayResponseLocations struct {
	ArrayResponse
	Data []*Location `json:"value"`
}

// ArrayResponseHistoricalLocations can be used to parse an array response containing historical locations
type ArrayResponseHistoricalLocations struct {
	ArrayResponse
	Data []*HistoricalLocation `json:"value"`
}

// ArrayResponseDatastreams can be used to parse an array response containing datastreams
type ArrayResponseDatastreams struct {
	ArrayResponse
	Data []*Datastream `json:"value"`
}

// ArrayResponseSensors can be used to parse an array response containing sensors
type ArrayResponseSensors struct {
	ArrayResponse
	Data []*Sensor `json:"value"`
}

// ArrayResponseObservedProperty can be used to parse an array response containing observed properties
type ArrayResponseObservedProperty struct {
	ArrayResponse
	Data []*ObservedProperty `json:"value"`
}

// ArrayResponseObservations can be used to parse an array response containing observations
type ArrayResponseObservations struct {
	ArrayResponse
	Data []*Observation `json:"value"`
}

// ArrayResponseFeaturesOfInterest can be used to parse an array response containing features of interest
type ArrayResponseFeaturesOfInterest struct {
	ArrayResponse
	Data []*FeatureOfInterest `json:"value"`
}
