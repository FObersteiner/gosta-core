package core

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCreateObservationsJson(t *testing.T) {
	// assert
	jsonString := "[{\"Datastream\": {\"@iot.id\": 1}, \"components\": [\"phenomenonTime\", \"result\", \"FeatureOfInterest/id\"],  		\"dataArray@iot.count\": 2, \"dataArray\": [[\"2010-12-23T10:20:00-0700\",20,1],[\"2010-12-23T10:21:00-0700\",30,1]]},{\"Datastream\": {\"@iot.id\": 2},\"components\": [\"phenomenonTime\",\"result\",\"FeatureOfInterest/id\"],\"dataArray@iot.count\": 1,\"dataArray\": [[\"2010-12-23T10:20:00-0700\",65,1]]}]"
	var obs []CreateObservations
	json.Unmarshal([]byte(jsonString), &obs)
	fmt.Printf("Datastreams: %v, DS1 Components: %s", len(obs), obs[0].Components)

	// 2 datastreams
	assert.Equal(t, 2, len(obs), "length of datastreams array should be 2")

	// first and 2ns datastream id's
	assert.Equal(t, float64(1), obs[0].Datastream.ID)
	assert.Equal(t, float64(2), obs[1].Datastream.ID)

	// 3 Components in first datastream
	assert.Equal(t, "phenomenonTime", obs[0].Components[0])
	assert.Equal(t, "result", obs[0].Components[1])
	assert.Equal(t, "FeatureOfInterest/id", obs[0].Components[2])

	// 2 results for datastream 1 in count field
	assert.Equal(t, int64(2), obs[0].Count)

	// 1 result for datastream 2 in count field
	assert.Equal(t, int64(1), obs[1].Count)

	// datastream 1 observation 1 values
	assert.Equal(t, "2010-12-23T10:20:00-0700", *obs[0].Data[0][0])
	assert.Equal(t, float64(20), *obs[0].Data[0][1])
	assert.Equal(t, float64(1), *obs[0].Data[0][2])

	// datastream 1 observation 2 values
	assert.Equal(t, "2010-12-23T10:21:00-0700", *obs[0].Data[1][0])
	assert.Equal(t, float64(30), *obs[0].Data[1][1])
	assert.Equal(t, float64(1), *obs[0].Data[1][2])

	// datastream 2 observation 2 values
	assert.Equal(t, "2010-12-23T10:20:00-0700", *obs[1].Data[0][0])
	assert.Equal(t, float64(65), *obs[1].Data[0][1])
	assert.Equal(t, float64(1), *obs[1].Data[0][2])
}
