package network

import (
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
)

func WriteError(w http.ResponseWriter, text string, code int) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(text))
}

func Write(w http.ResponseWriter, data interface{}) {
	_, _ = w.Write(marchy.Force(data))
}
