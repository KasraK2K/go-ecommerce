package pkg

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Name string `validate:"required"`
	Age  int    `validate:"min=0,max=130"`
}

func TestValidator(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		testData1 := TestStruct{Name: "John", Age: 30}
		expectedErrors1 := []*ValidationError(nil)
		response1 := Validator(testData1)
		if len(response1.Errors) != len(expectedErrors1) {
			t.Errorf("Validator() error count = %v; expected %v", len(response1.Errors), len(expectedErrors1))
		}
	})

	t.Run("Error", func(t *testing.T) {
		testData2 := TestStruct{Name: "", Age: 150}
		expectedErrors2 := []*ValidationError{
			{FailedField: "Name", Tag: "required", Value: ""},
			{FailedField: "Age", Tag: "max", Value: "130"},
		}
		response2 := Validator(testData2)
		if len(response2.Errors) != len(expectedErrors2) {
			t.Errorf("Validator() error count = %v; expected %v", len(response2.Errors), len(expectedErrors2))
		} else {
			for i, expected := range expectedErrors2 {
				actual := response2.Errors[i]
				// It should split because `StructNamespace` fill `FailedField` as fully qualified field name
				actualFieldName := strings.Split(actual.FailedField, ".")[1]
				expectedFieldName := strings.Split(expected.FailedField, " ")[0]

				if actualFieldName != expectedFieldName || actual.Tag != expected.Tag || actual.Value != expected.Value {
					t.Errorf("Validator() error at index %d; expected %v, but got %v", i, expected, actual)
				}
			}
		}
	})
}
