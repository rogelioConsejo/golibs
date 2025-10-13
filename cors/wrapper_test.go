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

// --- Helper functions for new tests ---

// getResponseHeaders executes a request with the given method and returns response headers.
func getResponseHeaders(w Wrapper, method string) http.Header {
	h := testHandler{}
	wrapped := w.Wrap(h)
	req := httptest.NewRequest(method, "/foo", nil)
	recorder := httptest.NewRecorder()
	wrapped.ServeHTTP(recorder, req)
	return recorder.Header()
}

// headerContainsValue checks if a comma-separated header contains the specified value (case-insensitive).
func headerContainsValue(headerValue, searchValue string) bool {
	values := strings.Split(headerValue, ",")
	for _, v := range values {
		if strings.EqualFold(strings.TrimSpace(v), searchValue) {
			return true
		}
	}
	return false
}

// --- New tests for Access-Control-Allow-Headers ---

func TestWrapper_AddHeader(t *testing.T) {
	t.Parallel()

	t.Run("Access-Control-Allow-Headers is empty by default on preflight requests", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()
		headers := getResponseHeaders(w, "OPTIONS")
		allowHeaders := headers.Get("Access-Control-Allow-Headers")

		if allowHeaders != "" {
			t.Errorf("Expected Access-Control-Allow-Headers to be empty, got: %s", allowHeaders)
		}
	})

	t.Run("AddHeader adds a single header to Access-Control-Allow-Headers", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()

		// Desired API: allow specifying a custom header
		w.AddHeader("Content-Type")

		headers := getResponseHeaders(w, "OPTIONS")
		allowHeaders := headers.Get("Access-Control-Allow-Headers")

		if !headerContainsValue(allowHeaders, "Content-Type") {
			t.Errorf("Expected Access-Control-Allow-Headers to contain 'Content-Type', got: %s", allowHeaders)
		}
	})

	t.Run("AddHeader adds multiple headers to Access-Control-Allow-Headers", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()

		// Desired API: allow specifying multiple custom headers
		w.AddHeader("Content-Type")
		w.AddHeader("Authorization")
		w.AddHeader("X-Custom-Header")

		headers := getResponseHeaders(w, "OPTIONS")
		allowHeaders := headers.Get("Access-Control-Allow-Headers")

		if !headerContainsValue(allowHeaders, "Content-Type") {
			t.Errorf("Expected Access-Control-Allow-Headers to contain 'Content-Type', got: %s", allowHeaders)
		}
		if !headerContainsValue(allowHeaders, "Authorization") {
			t.Errorf("Expected Access-Control-Allow-Headers to contain 'Authorization', got: %s", allowHeaders)
		}
		if !headerContainsValue(allowHeaders, "X-Custom-Header") {
			t.Errorf("Expected Access-Control-Allow-Headers to contain 'X-Custom-Header', got: %s", allowHeaders)
		}
	})

	t.Run("Access-Control-Allow-Headers only appears on OPTIONS requests", func(t *testing.T) {
		t.Parallel()
		w := NewWrapper()
		w.AddHeader("Content-Type")

		// Check that it appears on OPTIONS (preflight)
		optionsHeaders := getResponseHeaders(w, "OPTIONS")
		if optionsHeaders.Get("Access-Control-Allow-Headers") == "" {
			t.Error("Expected Access-Control-Allow-Headers to be set on OPTIONS request")
		}

		// Check that it doesn't appear on regular GET requests
		getHeaders := getResponseHeaders(w, "GET")
		if getHeaders.Get("Access-Control-Allow-Headers") != "" {
			t.Error("Expected Access-Control-Allow-Headers to be empty on GET request")
		}
	})
}
