package cors

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCORSEnabler(t *testing.T) {
	var h http.Handler = testHandler{}
	ceh := EnableCORS(h, http.MethodGet, http.MethodPost)
	req := httptest.NewRequest("OPTIONS", "/foo", nil)
	spy := httptest.NewRecorder()

	ceh.ServeHTTP(spy, req)
	if spy.Code != http.StatusNoContent {
		t.Error("Expected 204 OK")
	}
	if spy.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("Expected Access-Control-Allow-Origin header")
	}
	if !strings.Contains(spy.Header().Get("Access-Control-Allow-Methods"), http.MethodGet) {
		t.Error("Expected Access-Control-Allow-Methods header to contain GET")
	}
	if !strings.Contains(spy.Header().Get("Access-Control-Allow-Methods"), http.MethodPost) {
		t.Error("Expected Access-Control-Allow-Methods header to contain POST")
	}
	println(spy.Header().Get("Access-Control-Allow-Methods"))

	token := uuid.NewString()
	getReq := httptest.NewRequest("GET", fmt.Sprintf("/foo?foo=%s", token), nil)
	getSpy := httptest.NewRecorder()
	ceh.ServeHTTP(getSpy, getReq)
	if !strings.Contains(getSpy.Body.String(), "Hello World") {
		t.Error("Expected body to contain Hello World")
	}
	if !strings.Contains(getSpy.Body.String(), token) {
		t.Errorf("Expected body to contain %s: %s", token, getSpy.Body.String())
	}

}

type testHandler struct {
}

func (t testHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	foo := request.FormValue("foo")
	writer.Write([]byte(fmt.Sprintf("Hello World - %s", foo)))
}
