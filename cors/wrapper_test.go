package cors

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWrapper(t *testing.T) {
	t.Parallel()
	var w Wrapper = NewWrapper()
	if w == nil {
		t.Error("Expected non-nil Wrapper")
	}
}

func TestWrapper_Wrap(t *testing.T) {
	t.Parallel()
	t.Run("Wrap takes a handler and returns a handler", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()
		h := testHandler{}
		var wrapped http.Handler = w.Wrap(h)
		if wrapped == nil {
			t.Error("Expected non-nil http.Handler")
		}
	})
}

func TestWrapper_SetOrigin(t *testing.T) {
	t.Parallel()
	t.Run("setAllowedOrigin sets the Access-Control-Allow-Origin header", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()
		const url = "https://example.com"
		w.SetOrigin(url)
		h := testHandler{}
		wrapped := w.Wrap(h)
		req := httptest.NewRequest("GET", "/foo", nil)
		spy := httptest.NewRecorder()
		wrapped.ServeHTTP(spy, req)
		if spy.Header().Get("Access-Control-Allow-Origin") != url {
			t.Error("Expected Access-Control-Allow-Origin header to be set to", url)
		}
	})
	t.Run("Access-Control-Allow-Origin is set to * by default", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()
		h := testHandler{}
		wrapped := w.Wrap(h)
		req := httptest.NewRequest("GET", "/foo", nil)
		spy := httptest.NewRecorder()
		wrapped.ServeHTTP(spy, req)
		if spy.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Error("Expected Access-Control-Allow-Origin header to be set to *")
		}
	})
}

func TestWrapper_AddMethod(t *testing.T) {
	t.Parallel()
	t.Run("AddMethod adds a method to the Access-Control-Allow-Methods header", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()
		h := testHandler{}
		wrapped := w.Wrap(h)
		req := httptest.NewRequest("OPTIONS", "/foo", nil)
		spy := httptest.NewRecorder()
		wrapped.ServeHTTP(spy, req)
		if spy.Header().Get("Access-Control-Allow-Methods") != "OPTIONS" {
			t.Error("Expected Access-Control-Allow-Methods header to be set to OPTIONS")
		}
		w.AddMethod("GET")
		wrapped = w.Wrap(h)
		req = httptest.NewRequest("OPTIONS", "/foo", nil)
		spy = httptest.NewRecorder()
		wrapped.ServeHTTP(spy, req)
		methods := spy.Header().Get("Access-Control-Allow-Methods")
		hasOptions := strings.Contains(methods, "OPTIONS")
		hasGet := strings.Contains(methods, "GET")
		if !hasOptions || !hasGet {
			t.Error("Expected Access-Control-Allow-Methods header to contain OPTIONS and GET")
		}
	})
}

func TestWrapper_SetAllowCredentials(t *testing.T) {
	t.Parallel()
	w := NewWrapper()
	h := testHandler{}
	t.Run("Access-Control-Allow-Credentials is not set by default", func(t *testing.T) {
		wrapped := w.Wrap(h)
		req := httptest.NewRequest("GET", "/foo", nil)
		spy := httptest.NewRecorder()
		wrapped.ServeHTTP(spy, req)
		if spy.Header().Get("Access-Control-Allow-Credentials") != "" {
			t.Error("Expected Access-Control-Allow-Credentials header to be empty")
		}
	})
	t.Run("SetAllowCredentials sets the Access-Control-Allow-Credentials header", func(t *testing.T) {
		w.SetAllowCredentials()
		wrapped := w.Wrap(h)
		req := httptest.NewRequest("GET", "/foo", nil)
		spy := httptest.NewRecorder()
		wrapped.ServeHTTP(spy, req)
		if spy.Header().Get("Access-Control-Allow-Credentials") != "true" {
			t.Error("Expected Access-Control-Allow-Credentials header to be set to true")
		}
	})
}
