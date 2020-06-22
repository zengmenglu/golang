package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBinding(t *testing.T) {
	data := `{"message":"aha binding","nick":"aha nil"}`
	route := setFormRoute()
	rsp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/form_post", strings.NewReader(data))
	route.ServeHTTP(rsp, req)
	if rsp.Code != http.StatusOK {
		t.Error("fail")
	}
	t.Logf("rsp:%+v", rsp)
}
