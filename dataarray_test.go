package core

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jsonString = "[{\"Datastream\": {\"@iot.id\": 1}, \"components\": [\"phenomenonTime\", \"result\", \"FeatureOfInterest/id\"], 	\"dataArray@iot.count\": 2, \"dataArray\": [[\"2010-12-23T10:20:00-0700\",20,1],	[\"2010-12-23T10:21:00-0700\",30.1,1]]},	{\"Datastream\": {\"@iot.id\": 2},	\"components\": [\"phenomenonTime\", \"result\", \"validTime\", \"resultTime\", \"resultQuality\", \"parameters\", \"FeatureOfInterest/id\"],	\"dataArray@iot.count\": 1,	\"dataArray\": [[\"2010-12-23T10:20:00-0700\",\"soep\",\"2011-12-23T10:20:00-0700\",\"2012-12-23T10:20:00-0700\",\"goed\", {\"test1\": \"soep\", \"test2\": \"ballen\"},2]]}]"
)

func TestParseCreateObservations(t *testing.T) {
	//arrange
	co := &CreateObservations{}

	//act
	err := co.ParseEntity([]byte(jsonString))

	//assert
	assert.Nil(t, err, "Parsing JSON should not give an error")

	// 2 datastreams
	assert.Equal(t, 2, len(co.Datastreams), "length of datastreams array should be 2")

	// first and 2ns datastream id's
	assert.Equal(t, float64(1), co.Datastreams[0].ID)
	assert.Equal(t, float64(2), co.Datastreams[1].ID)

	// 2 observations for datastream 1
	assert.Equal(t, int(2), len(co.Datastreams[0].Observations))

	// 1 observation for datastream 2
	assert.Equal(t, int(1), len(co.Datastreams[1].Observations))

	// values for DS 1 Observation 1
	var res1 int
	json.Unmarshal(co.Datastreams[0].Observations[0].Result, &res1)
	assert.Equal(t, "2010-12-23T10:20:00-0700", co.Datastreams[0].Observations[0].PhenomenonTime)
	assert.Equal(t, 20, res1)
	assert.Equal(t, float64(1), co.Datastreams[0].Observations[0].FeatureOfInterest.ID)

	// values for DS 1 Observation 2
	var res2 float64
	json.Unmarshal(co.Datastreams[0].Observations[1].Result, &res2)
	assert.Equal(t, "2010-12-23T10:21:00-0700", co.Datastreams[0].Observations[1].PhenomenonTime)
	assert.Equal(t, 30.1, res2)
	assert.Equal(t, float64(1), co.Datastreams[0].Observations[1].FeatureOfInterest.ID)

	// values for DS 2 Observation 1
	var res3 string
	json.Unmarshal(co.Datastreams[1].Observations[0].Result, &res3)
	assert.Equal(t, "2010-12-23T10:20:00-0700", co.Datastreams[1].Observations[0].PhenomenonTime)
	assert.Equal(t, "soep", res3)
	assert.Equal(t, "2011-12-23T10:20:00-0700", co.Datastreams[1].Observations[0].ValidTime)
	assert.Equal(t, "2012-12-23T10:20:00-0700", *co.Datastreams[1].Observations[0].ResultTime)
	assert.Equal(t, "goed", co.Datastreams[1].Observations[0].ResultQuality)
	assert.Equal(t, "soep", co.Datastreams[1].Observations[0].Parameters["test1"])
	assert.Equal(t, "ballen", co.Datastreams[1].Observations[0].Parameters["test2"])
	assert.Equal(t, float64(2), co.Datastreams[1].Observations[0].FeatureOfInterest.ID)
}
