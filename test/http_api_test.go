package test

import (
	"integrtest/startup"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestHTTPApi(t *testing.T) {
	httpServer := startup.InitializeHTTPServer()
	defer func() {
		if err := httpServer.Shutdown(); err != nil {
			t.Error("test server shutdown error:", err)
		}
	}()

	go func() {
		if err := httpServer.Listen(":3123"); err != nil {
			t.Error("test server listening error:", err)
		}
	}()

	e := httpexpect.Default(t, "http://localhost:3123")

	e.GET("/").Expect().Status(http.StatusOK)
	// TODO: place tests here
}
