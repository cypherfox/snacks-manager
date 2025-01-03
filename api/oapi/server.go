//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config-server.yaml snacks-manager-1.0.0.oapi-3.0.3.yml

package oapi

import (
	"context"
	"net/http"

	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

const (
	HEADER_KEY_AUTH_USER = "X-Snackmgr-Authenticated-User"
	OWNER_KEY            = "owner"
)

type ApiServer struct {
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

func NewApiHandler() ServerInterface {
	server := NewApiServer()
	handler := NewStrictHandler(server,
		[]strictnethttp.StrictHTTPMiddlewareFunc{ProcessAuthHeader})

	return handler
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

// GetSnacks implements StrictServerInterface.
func (a *ApiServer) GetSnacks(ctx context.Context, request GetSnacksRequestObject) (GetSnacksResponseObject, error) {
	panic("unimplemented")
}

// GetSnacksSnackId implements StrictServerInterface.
func (a *ApiServer) GetSnacksSnackId(ctx context.Context, request GetSnacksSnackIdRequestObject) (GetSnacksSnackIdResponseObject, error) {
	panic("unimplemented")
}

// GetTest implements StrictServerInterface.
func (a *ApiServer) GetTest(ctx context.Context, request GetTestRequestObject) (GetTestResponseObject, error) {
	panic("unimplemented")
}

// PostPurchaseAcknowledge implements StrictServerInterface.
func (a *ApiServer) PostPurchaseAcknowledge(ctx context.Context, request PostPurchaseAcknowledgeRequestObject) (PostPurchaseAcknowledgeResponseObject, error) {
	panic("unimplemented")
}

// PostPurchaseOrder implements StrictServerInterface.
func (a *ApiServer) PostPurchaseOrder(ctx context.Context, request PostPurchaseOrderRequestObject) (PostPurchaseOrderResponseObject, error) {
	panic("unimplemented")
}

// PostPurchaseProcessOrderId implements StrictServerInterface.
func (a *ApiServer) PostPurchaseProcessOrderId(ctx context.Context, request PostPurchaseProcessOrderIdRequestObject) (PostPurchaseProcessOrderIdResponseObject, error) {
	panic("unimplemented")
}
