// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// health go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/client/fetcher/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/http"
)

// MountShowHandler configures the mux to serve the "health" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/health", f)
}
