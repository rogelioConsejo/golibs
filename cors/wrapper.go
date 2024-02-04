// Package cors provides a middleware that easily enables CORS for existing handlers.
package cors

import "net/http"

// NewWrapper returns a new Wrapper. The implementation is abstracted and encapsulated in the returned value.
func NewWrapper() Wrapper {
	return &wrapper{}
}

// Wrapper is a type that can wrap a http.Handler to enable CORS
// It allows to set the origin, the methods and if credentials are allowed
// The wrap configuration must be done before calling Wrap
type Wrapper interface {
	Wrap(http.Handler) http.Handler
	SetOrigin(URL)
	AddMethod(s Method)
	SetAllowCredentials()
}

// URL is a type that represents a URL
type URL string

// Method is a type that represents an HTTP method
type Method string

type wrapper struct {
	origin            URL
	methods           map[Method]bool
	allowsCredentials bool
}

// SetAllowCredentials sets the CORS configuration to allow credentials
func (w *wrapper) SetAllowCredentials() {
	w.allowsCredentials = true
}

// AddMethod adds a Method to the CORS wrapper configuration
func (w *wrapper) AddMethod(m Method) {
	if w.methods == nil {
		w.methods = make(map[Method]bool)
	}
	w.methods[m] = true
}

// SetOrigin sets the origin for the CORS configuration to the given URL, otherwise it will default to *
func (w *wrapper) SetOrigin(o URL) {
	w.origin = o
}

// Wrap returns a new http.Handler that wraps the given handler and enables CORS with the given configuration
func (w *wrapper) Wrap(h http.Handler) http.Handler {
	var methods []string
	for m := range w.methods {
		methods = append(methods, string(m))
	}
	return corsEnabledHandler{handler: h, origin: w.origin, methods: methods, allowsCredentials: w.allowsCredentials}
}
