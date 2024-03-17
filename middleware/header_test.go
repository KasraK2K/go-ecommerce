package middleware

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"app/config"
)

func TestAddHeaderMiddleware(t *testing.T) {
	config.SetConfig()

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	app.Use(AddHeaderMiddleware)

	req := &http.Request{
		Method: http.MethodPost,
		URL:    &url.URL{Path: "/"},
		Header: http.Header{},
	}
	resp, err := app.Test(req)
	if err != nil {
		t.Errorf("AddHeaderMiddleware error on sending test request")
	}

	allHeaders := resp.Header
	for headerName, headerValue := range allHeaders {
		if headerValue[0] == "" {
			t.Errorf("AddHeaderMiddleware() did not set the %s header", headerName)
		}
	}
}
