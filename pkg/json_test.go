package pkg

import (
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"app/config"
)

type TestJsonStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMarshal(t *testing.T) {
	mockData := TestJsonStruct{
		Name: "Kasra",
		Age:  37,
	}

	t.Run("Development", func(t *testing.T) {
		config.AppConfig.MODE = "development"
		expectedData := mockData

		actualResult, err := Marshal(expectedData)
		if err != nil {
			t.Errorf("It has error on Marshal: %v", err)
		}
		expectedResult, _ := json.MarshalIndent(expectedData, "", "  ")

		if string(actualResult) != string(expectedResult) {
			t.Errorf("Actual result is not same as expected result")
		}
	})

	t.Run("Production", func(t *testing.T) {
		config.AppConfig.MODE = "production"
		expectedData := mockData

		actualResult, err := Marshal(expectedData)
		if err != nil {
			t.Errorf("It has error on Marshal: %v", err)
		}
		expectedResult, _ := json.Marshal(expectedData)

		if string(actualResult) != string(expectedResult) {
			t.Errorf("Actual result is not same as expected result")
		}
	})
}

func TestJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockData := TestJsonStruct{
			Name: "Kasra",
			Age:  37,
		}
		app := fiber.New()
		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		err := JSON(c, mockData, http.StatusOK)
		if err != nil {
			t.Errorf("error on JSON command: %v", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		mockData := make(chan int)
		app := fiber.New()
		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		err := JSON(c, mockData, http.StatusInternalServerError)
		if err == nil {
			t.Errorf("The error is nil but we expected to be something")
		}
	})
}
