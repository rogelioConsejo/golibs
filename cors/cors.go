package cors

import (
	"net/http"
	"strings"
)

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
	c.SetAllowedOrigin(writer)
	if request.Method == http.MethodOptions {
		c.HandleOptionsResponse(writer)
		return
	}
	c.handler.ServeHTTP(writer, request)
}

func (c corsEnabledHandler) HandleOptionsResponse(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	methods := append(c.methods, http.MethodOptions)
	writer.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	if c.allowsCredentials {
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	writer.WriteHeader(http.StatusOK)
}

func (c corsEnabledHandler) SetAllowedOrigin(writer http.ResponseWriter) {
	const defaultOrigin = "*"
	var origin string
	if c.origin != "" {
		origin = string(c.origin)
	} else {
		origin = defaultOrigin
	}
	writer.Header().Set("Access-Control-Allow-Origin", origin)
}
