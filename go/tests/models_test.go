package tests

import (
	"cloud.google.com/go/civil"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/specgen-io/test-models-go/test_service_models"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
)

func TestMessageFields(t *testing.T) {
	data := test_service_models.Message{
		0,
	}

	jsonStr := `{"field":0}`

	var actualData test_service_models.Message
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNestedFields(t *testing.T) {
	data := test_service_models.Parent{
		"the string",
		test_service_models.Nested{
			"the nested string",
		},
	}

	jsonStr := `{"field":"the string","nested":{"field":"the nested string"}}`

	var actualData test_service_models.Parent
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNumericFields(t *testing.T) {
	decimalField, _ := decimal.NewFromString("1.23")
	data := test_service_models.NumericFields{
		0,
		0,
		1.23,
		1.23,
		decimalField,
	}

	jsonStr := `{"int_field":0,"long_field":0,"float_field":1.23,"double_field":1.23,"decimal_field":"1.23"}`

	var actualData test_service_models.NumericFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNumericFieldsNegative(t *testing.T) {
	jsonStr := `{"int_field":0,"long_field":0,"float_field":1.23,"double_field":1.23,"decimal_field":nil}`

	var actualData test_service_models.NumericFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)

	assert.ErrorContains(t, err, "null")
}

func TestNonNumericFields(t *testing.T) {
	dateField, _ := civil.ParseDate("2019-11-30")
	timeDateField, _ := civil.ParseDateTime("2019-11-30T17:45:55")
	data := test_service_models.NonNumericFields{
		true,
		"the string",
		uuid.MustParse("123e4567-e89b-12d3-a456-426655440000"),
		dateField,
		timeDateField,
	}

	jsonStr := `{"boolean_field":true,"string_field":"the string","uuid_field":"123e4567-e89b-12d3-a456-426655440000","date_field":"2019-11-30","datetime_field":"2019-11-30T17:45:55"}`

	var actualData test_service_models.NonNumericFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestNonNumericFieldsNegative(t *testing.T) {
	jsonStr := `{"boolean_field":true,"string_field":"the string","uuid_field":nil,"date_field":"2019-11-30","datetime_field":"2019-11-30T17:45:55"}`

	var actualData test_service_models.NonNumericFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)

	assert.ErrorContains(t, err, "null")
}

func TestArrayFields(t *testing.T) {
	data := test_service_models.ArrayFields{
		[]int{1, 2, 3},
		[]string{"one", "two", "three"},
	}

	jsonStr := `{"int_array_field":[1,2,3],"string_array_field":["one","two","three"]}`

	var actualData test_service_models.ArrayFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestArrayFieldsNegative(t *testing.T) {
	jsonStr := `{"int_array_field":[nil],"string_array_field":["one","two","three"]}`

	var actualData test_service_models.ArrayFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)

	assert.ErrorContains(t, err, "null")
}

func TestMapFields(t *testing.T) {
	data := test_service_models.MapFields{
		map[string]int{"one": 1, "two": 2},
		map[string]string{"one": "first", "two": "second"},
	}

	jsonStr := `{"int_map_field":{"one":1,"two":2},"string_map_field":{"one":"first","two":"second"}}`

	var actualData test_service_models.MapFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestMapFieldsNegative(t *testing.T) {
	jsonStr := `{"int_map_field":{"one":1,"two":2},"string_map_field":{"one":nil,"two":"second"}}`

	var actualData test_service_models.MapFields
	err := json.Unmarshal([]byte(jsonStr), &actualData)

	assert.ErrorContains(t, err, "null")
}

func TestOptionalFields(t *testing.T) {
	var optionalInt int = 123
	var optionalString = "the string"
	data := test_service_models.OptionalFields{
		&optionalInt,
		&optionalString,
	}

	jsonStr := `{"int_option_field":123,"string_option_field":"the string"}`

	var actualData test_service_models.OptionalFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestOptionalNil(t *testing.T) {
	data := test_service_models.OptionalFields{
		nil,
		nil,
	}

	jsonStr := `{"int_option_field":null,"string_option_field":null}`

	var actualData test_service_models.OptionalFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestRawJsonField(t *testing.T) {
	data := test_service_models.RawJsonField{
		json.RawMessage(`{"the_array":[1,"some string"],"the_object":{"the_bool":true,"the_string":"some value"},"the_scalar":123}`),
	}

	jsonStr := `{"json_field":{"the_array":[1,"some string"],"the_object":{"the_bool":true,"the_string":"some value"},"the_scalar":123}}`

	var actualData test_service_models.RawJsonField
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestRawJsonFieldNegative(t *testing.T) {
	jsonStr := `{"json_field":{"the_array":[1,"some string"],"the_object":{"the_bool":nil,"the_string":"some value"},"the_scalar":123}}`

	var actualData test_service_models.RawJsonField
	err := json.Unmarshal([]byte(jsonStr), &actualData)

	assert.ErrorContains(t, err, "null")
}

func TestOrderCreated(t *testing.T) {
	data := test_service_models.OrderCreated{
		uuid.MustParse("58d5e212-165b-4ca0-909b-c86b9cee0111"),
		"SNI/01/136/0500",
		3,
	}

	jsonStr := `{"id":"58d5e212-165b-4ca0-909b-c86b9cee0111","sku":"SNI/01/136/0500","quantity":3}`

	var actualData test_service_models.OrderCreated
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestOrderCreatedNegative(t *testing.T) {
	jsonStr := `{"id":"58d5e212-165b-4ca0-909b-c86b9cee0111","sku":"SNI/01/136/0500","quantity":nil}`

	var actualData test_service_models.OrderCreated
	err := json.Unmarshal([]byte(jsonStr), &actualData)

	assert.ErrorContains(t, err, "null")
}

func TestEnumFields(t *testing.T) {
	data := test_service_models.EnumFields{
		test_service_models.ChoiceSecondChoice,
	}

	jsonStr := `{"enum_field":"SECOND_CHOICE"}`

	var actualData test_service_models.EnumFields
	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}

func TestEnumFieldsNegative(t *testing.T) {
	jsonStr := `{"enum_field":"gfgnfg"}`

	var actualData test_service_models.EnumFields

	err := json.Unmarshal([]byte(jsonStr), &actualData)
	assert.ErrorContains(t, err, "gfgnfg")
}

func TestOneOfFields(t *testing.T) {
	data := test_service_models.OrderEvent{
		nil,
		&test_service_models.OrderChanged{uuid.MustParse("58d5e212-165b-4ca0-909b-c86b9cee0111"), 3},
		nil,
	}

	jsonStr := `{"changed":{"id":"58d5e212-165b-4ca0-909b-c86b9cee0111","quantity":3}}`

	var actualData test_service_models.OrderEvent

	actualJson, err := json.Marshal(data)
	assert.NilError(t, err)

	err = json.Unmarshal([]byte(jsonStr), &actualData)
	assert.NilError(t, err)

	assert.Equal(t, jsonStr, string(actualJson))
	assert.Equal(t, reflect.DeepEqual(data, actualData), true)
}
