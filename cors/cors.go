package cors

import (
	"net/http"
	"strings"
)

func EnableCORS(h http.Handler, methods ...string) http.Handler {
	return corsEnabledHandler{handler: h, methods: methods}
}

type corsEnabledHandler struct {
	handler http.Handler
	methods []string
}

func (c corsEnabledHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusNoContent)
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", strings.Join(c.methods, ","))
		writer.WriteHeader(http.StatusOK)
		return
	}
	c.handler.ServeHTTP(writer, request)
}
