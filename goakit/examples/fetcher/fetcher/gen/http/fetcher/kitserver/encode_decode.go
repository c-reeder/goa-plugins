// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// fetcher go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/client/fetcher/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/http"
	"goa.design/plugins/goakit/examples/client/fetcher/gen/http/fetcher/server"
)

// EncodeFetchResponse returns a go-kit EncodeResponseFunc suitable for
// encoding fetcher fetch responses.
func EncodeFetchResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeFetchResponse(encoder)
}

// DecodeFetchRequest returns a go-kit DecodeRequestFunc suitable for decoding
// fetcher fetch requests.
func DecodeFetchRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeFetchRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeFetchResponse returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the fetcher fetch endpoint.
func EncodeFetchError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	enc := server.EncodeFetchError(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		enc(ctx, w, v.(error))
		return nil
	}
}
