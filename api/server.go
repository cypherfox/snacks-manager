//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config-server.yaml snacks-manager-1.0.0.oapi-3.0.3.yml

package oapi

import (
	"context"
	"net/http"

	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

const (
	HEADER_KEY_AUTH_USER = "X-Snackmgr-Authenticated-User"
)

type ApiServer struct {
	model *model.DataModel
}

var _ StrictServerInterface = (*ApiServer)(nil)

type RequestHeaderKey string

/*
ProcessAuthHeader is a middleware to transfer the authentication header "X-Shmits-Authenticated-User" into the context for
the call to the Strict Server Interface.

	Since the requirement for the existence of a valid user depends on the actual method an path being accessed, validation
	is handled in the individual methods of the Strict Service Interface implementation.
*/
func ProcessAuthHeader(f StrictHandlerFunc, _ string) StrictHandlerFunc {

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		var newCtx context.Context

		// check if header is set at all.
		_, ok := r.Header[HEADER_KEY_AUTH_USER]
		if ok {
			// this has more compliant processing for edge cases like multiple values and
			// case insensitive matches
			owner := r.Header.Get(HEADER_KEY_AUTH_USER)

			newCtx = context.WithValue(ctx, OWNER_KEY, owner)

		} else {
			newCtx = ctx
		}

		return f(newCtx, w, r, request)
	}
}

func NewApiHandler(m *model.EnterpriseModel) ServerInterface {
	server := NewApiServer(m)
	handler := NewStrictHandler(server,
		[]strictnethttp.StrictHTTPMiddlewareFunc{ProcessAuthHeader})

	return handler
}

func NewApiServer(m *model.EnterpriseModel) *ApiServer {
	return &ApiServer{
		model: m,
	}
}
