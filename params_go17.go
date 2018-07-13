// +build go1.7

package pathrouter

import (
	"context"
)

type paramsKey struct{}

// ParamsKey is the request context key under which URL params are stored.
//
// This is only present from go 1.7.
var ParamsKey = paramsKey{}

// ParamsFromContext pulls the URL parameters from a request context,
// or returns nil if none are present.
//
// This is only present from go 1.7.
func ParamsFromContext(ctx context.Context) Params {
	p, _ := ctx.Value(ParamsKey).(Params)
	return p
}
