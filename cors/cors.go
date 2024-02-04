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
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	if request.Method == http.MethodOptions {
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		methods := append(c.methods, http.MethodOptions)
		writer.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.WriteHeader(http.StatusOK)
		return
	}
	c.handler.ServeHTTP(writer, request)
}
