package tests

import (
	"encoding/json"
	"github.com/specgen-io/test-models-go/spec_v2"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestMessageV2(t *testing.T) {
	data := spec_v2.Message{
		"the string",
	}

	jsonStr := `{"field":"the string"}`

	var actualData spec_v2.Message
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}
