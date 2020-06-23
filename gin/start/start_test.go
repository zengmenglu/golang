package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	route := setRoute()
	rsp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	route.ServeHTTP(rsp, req)
	if rsp.Code != http.StatusOK || rsp.Body.String() != "{\"message\":\"pong\"}" {
		t.Error("fail")
	}
}
