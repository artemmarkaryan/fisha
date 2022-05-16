package network

import (
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
)

func InternalError(w http.ResponseWriter) {
	w.WriteHeader(500)
}

func WriteBadRequestError(w http.ResponseWriter, text string) {
	w.WriteHeader(400)
	_, _ = w.Write([]byte("bad request: " + text))
}

func WriteError(w http.ResponseWriter, text string, code int) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(text))
}

func Write(w http.ResponseWriter, data interface{}) {
	_, _ = w.Write(marchy.Force(data))
}
