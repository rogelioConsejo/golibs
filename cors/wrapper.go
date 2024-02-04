package cors

import "net/http"

func NewWrapper() Wrapper {
	return &wrapper{}
}

type Wrapper interface {
	Wrap(http.Handler) http.Handler
	SetOrigin(URL)
	AddMethod(s Method)
	SetAllowCredentials()
}

type URL string

type Method string

type wrapper struct {
	origin            URL
	methods           map[Method]bool
	allowsCredentials bool
}

func (w *wrapper) SetAllowCredentials() {
	w.allowsCredentials = true
}

func (w *wrapper) AddMethod(m Method) {
	if w.methods == nil {
		w.methods = make(map[Method]bool)
	}
	w.methods[m] = true
}

func (w *wrapper) SetOrigin(o URL) {
	w.origin = o
}

func (w *wrapper) Wrap(h http.Handler) http.Handler {
	var methods []string
	for m := range w.methods {
		methods = append(methods, string(m))
	}
	return corsEnabledHandler{handler: h, origin: w.origin, methods: methods, allowsCredentials: w.allowsCredentials}
}
