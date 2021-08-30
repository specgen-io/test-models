package tests

import (
	"encoding/json"
	"gotest.tools/v3/assert"
	"reflect"
	"test-models/test_service_models/v2"
	"testing"
)

func TestMessageV2(t *testing.T) {
	data := v2.Message{
		"the string",
	}

	jsonStr := `{"field":"the string"}`

	var actualData v2.Message
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}
