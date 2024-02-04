package cors

import (
	"net/http"
	"strings"
)

// EnableCORS is a middleware that enables CORS for the given handler
// It is a basic implementation of CORS that allows all origins and all methods
func EnableCORS(h http.Handler, methods ...string) http.Handler {
	return corsEnabledHandler{handler: h, methods: methods}
}

type corsEnabledHandler struct {
	handler           http.Handler
	methods           []string
	origin            URL
	allowsCredentials bool
}

func (c corsEnabledHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c.setAllowedOrigin(writer)
	if request.Method == http.MethodOptions {
		c.handleOptionsResponse(writer)
		return
	}
	c.handler.ServeHTTP(writer, request)
}

func (c corsEnabledHandler) handleOptionsResponse(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	methods := append(c.methods, http.MethodOptions)
	writer.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	if c.allowsCredentials {
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	writer.WriteHeader(http.StatusOK)
}

func (c corsEnabledHandler) setAllowedOrigin(writer http.ResponseWriter) {
	var origin string
	if c.origin != "" {
		origin = string(c.origin)
	} else {
		origin = DefaultOrigin
	}
	writer.Header().Set("Access-Control-Allow-Origin", origin)
}

const DefaultOrigin = "*"
