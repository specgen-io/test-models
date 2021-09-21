package tests

import (
	"cloud.google.com/go/civil"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gotest.tools/v3/assert"
	"reflect"
	"test-models/spec/models"
	"testing"
)

func TestMessageFields(t *testing.T) {
	data := models.Message{
		0,
	}

	jsonStr := `{"field":0}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.Message
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNestedFields(t *testing.T) {
	data := models.Parent{
		"the string",
		models.Nested{
			"the nested string",
		},
	}

	jsonStr := `{"field":"the string","nested":{"field":"the nested string"}}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.Parent
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNumericFields(t *testing.T) {
	decimalField, _ := decimal.NewFromString("1.23")
	data := models.NumericFields{
		0,
		0,
		1.23,
		1.23,
		decimalField,
	}

	jsonStr := `{"int_field":0,"long_field":0,"float_field":1.23,"double_field":1.23,"decimal_field":"1.23"}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.NumericFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNumericFieldsNegative(t *testing.T) {
	jsonStr := `{"int_field":0,"long_field":0,"float_field":1.23,"double_field":1.23,"decimal_field":nil}`

	var actualData models.NumericFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "null")
}

func TestNonNumericFields(t *testing.T) {
	dateField, _ := civil.ParseDate("2019-11-30")
	timeDateField, _ := civil.ParseDateTime("2019-11-30T17:45:55")
	data := models.NonNumericFields{
		true,
		"the string",
		uuid.MustParse("123e4567-e89b-12d3-a456-426655440000"),
		dateField,
		timeDateField,
	}

	jsonStr := `{"boolean_field":true,"string_field":"the string","uuid_field":"123e4567-e89b-12d3-a456-426655440000","date_field":"2019-11-30","datetime_field":"2019-11-30T17:45:55"}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.NonNumericFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNonNumericFieldsNegative(t *testing.T) {
	jsonStr := `{"boolean_field":true,"string_field":"the string","uuid_field":nil,"date_field":"2019-11-30","datetime_field":"2019-11-30T17:45:55"}`

	var actualData models.NonNumericFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "null")
}

func TestArrayFields(t *testing.T) {
	data := models.ArrayFields{
		[]int{1, 2, 3},
		[]string{"one", "two", "three"},
	}

	jsonStr := `{"int_array_field":[1,2,3],"string_array_field":["one","two","three"]}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.ArrayFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestArrayFieldsNegative(t *testing.T) {
	jsonStr := `{"int_array_field":[nil],"string_array_field":["one","two","three"]}`

	var actualData models.ArrayFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "null")
}

func TestMapFields(t *testing.T) {
	data := models.MapFields{
		map[string]int{"one": 1, "two": 2},
		map[string]string{"one": "first", "two": "second"},
	}

	jsonStr := `{"int_map_field":{"one":1,"two":2},"string_map_field":{"one":"first","two":"second"}}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.MapFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestMapFieldsNegative(t *testing.T) {
	jsonStr := `{"int_map_field":{"one":1,"two":2},"string_map_field":{"one":nil,"two":"second"}}`

	var actualData models.MapFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "null")
}

func TestOptionalFields(t *testing.T) {
	var optionalInt = 123
	var optionalString = "the string"
	data := models.OptionalFields{
		&optionalInt,
		&optionalString,
	}

	jsonStr := `{"int_option_field":123,"string_option_field":"the string"}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.OptionalFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestOptionalNil(t *testing.T) {
	data := models.OptionalFields{
		nil,
		nil,
	}

	jsonStr := `{"int_option_field":null,"string_option_field":null}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.OptionalFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestRawJsonField(t *testing.T) {
	data := models.RawJsonField{
		json.RawMessage(`{"the_array":[1,"some string"],"the_object":{"the_bool":true,"the_string":"some value"},"the_scalar":123}`),
	}

	jsonStr := `{"json_field":{"the_array":[1,"some string"],"the_object":{"the_bool":true,"the_string":"some value"},"the_scalar":123}}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.RawJsonField
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestRawJsonFieldNegative(t *testing.T) {
	jsonStr := `{"json_field":{"the_array":[1,"some string"],"the_object":{"the_bool":nil,"the_string":"some value"},"the_scalar":123}}`

	var actualData models.RawJsonField
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "null")
}

func TestOrderCreated(t *testing.T) {
	data := models.OrderCreated{
		uuid.MustParse("58d5e212-165b-4ca0-909b-c86b9cee0111"),
		"SNI/01/136/0500",
		3,
	}

	jsonStr := `{"id":"58d5e212-165b-4ca0-909b-c86b9cee0111","sku":"SNI/01/136/0500","quantity":3}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.OrderCreated
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestOrderCreatedNegative(t *testing.T) {
	jsonStr := `{"id":"58d5e212-165b-4ca0-909b-c86b9cee0111","sku":"SNI/01/136/0500","quantity":nil}`

	var actualData models.OrderCreated
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "null")
}

func TestEnumFields(t *testing.T) {
	data := models.EnumFields{
		models.ChoiceSecondChoice,
	}

	jsonStr := `{"enum_field":"SECOND_CHOICE"}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.EnumFields
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestEnumFieldsNegative(t *testing.T) {
	jsonStr := `{"enum_field":"gfgnfg"}`

	var actualData models.EnumFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "gfgnfg")
}

func TestOneOfFields(t *testing.T) {
	data := models.OrderEvent{
		Changed: &models.OrderChanged{uuid.MustParse("58d5e212-165b-4ca0-909b-c86b9cee0111"), 3},
	}

	jsonStr := `{"changed":{"id":"58d5e212-165b-4ca0-909b-c86b9cee0111","quantity":3}}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.OrderEvent
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestDiscriminatedOneOfFields(t *testing.T) {
	data := models.OrderEventDiscriminated{
		Changed: &models.OrderChanged{uuid.MustParse("58d5e212-165b-4ca0-909b-c86b9cee0111"), 3},
	}

	jsonStr := `{"_type":"changed","id":"58d5e212-165b-4ca0-909b-c86b9cee0111","quantity":3}`

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)
	assert.Equal(t, jsonStr, string(actualJson))

	var actualData models.OrderEventDiscriminated
	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}
