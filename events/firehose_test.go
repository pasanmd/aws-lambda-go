// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
package events

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events/test"
	"github.com/stretchr/testify/assert"
)

func TestFirehoseEventMarshaling(t *testing.T) {
	testMarshaling(t, &KinesisFirehoseEvent{}, "./testdata/kinesis-firehose-event.json")
}

func TestFirehoseResponseMarshaling(t *testing.T) {
	testMarshaling(t, &KinesisFirehoseResponse{}, "./testdata/kinesis-firehose-response.json")
}

func testMarshaling(t *testing.T, inputEvent interface{}, jsonFile string) {
	// 1. read JSON from file
	inputJSON := test.ReadJSONFromFile(t, jsonFile)

	// 2. de-serialize into Go object
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// 3. serialize to JSON
	outputJSON, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	// 4. check result
	assert.JSONEq(t, string(inputJSON), string(outputJSON))
}

func TestSampleTransformation(t *testing.T) {
	inputJSON := test.ReadJSONFromFile(t, "./testdata/kinesis-firehose-event.json")

	// de-serialize into Go object
	var inputEvent KinesisFirehoseEvent
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	response := toUpperHandler(context.TODO(), inputEvent)

	inputString := string(inputEvent.Records[0].Data)
	expectedString := strings.ToUpper(inputString)
	actualString := string(response.Records[0].Data)
	assert.Equal(t, actualString, expectedString)
}

func toUpperHandler(ctx context.Context, evnt KinesisFirehoseEvent) KinesisFirehoseResponse {
	var response KinesisFirehoseResponse

	for _, record := range evnt.Records {
		// Transform data: ToUpper the data
		var transformedRecord KinesisFirehoseResponseRecord
		transformedRecord.RecordID = record.RecordID
		transformedRecord.Result = KinesisFirehoseTransformedStateOk
		transformedRecord.Data = []byte(strings.ToUpper(string(record.Data)))

		response.Records = append(response.Records, transformedRecord)
	}

	return response
}

func TestKinesisFirehoseMarshalingMalformedJson(t *testing.T) {
	test.TestMalformedJson(t, KinesisFirehoseEvent{})
}
