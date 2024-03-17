package pkg

import (
	"net/http"
	"testing"
)

func TestAddMetaData(t *testing.T) {
	mockData := "Test data"

	t.Run("Success", func(t *testing.T) {
		expectedSuccess := true
		expectedResult := mockData
		md := AddMetaData(mockData, http.StatusOK)
		if md.SUCCESS != expectedSuccess {
			t.Errorf("AddMetaData() success flag = %v; expected %v", md.SUCCESS, expectedSuccess)
		}
		if md.RESULT != expectedResult {
			t.Errorf("AddMetaData() result = %v; expected %v", md.RESULT, expectedResult)
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedSuccess := false
		expectedError := mockData
		md := AddMetaData(mockData, http.StatusInternalServerError)
		if md.SUCCESS != expectedSuccess {
			t.Errorf("AddMetaData() success flag = %v; expected %v", md.SUCCESS, expectedSuccess)
		}
		if md.ERRORS != expectedError {
			t.Errorf("AddMetaData() error = %v; expected %v", md.ERRORS, expectedError)
		}
	})
}
