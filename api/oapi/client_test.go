package oapi_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cypherfox/snacks-manager/api/oapi"
	"github.com/cypherfox/snacks-manager/internal/backend"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestPurchase(t *testing.T) {
	server, err := newListener()
	require.Nil(t, err)

	defer server.Shutdown(context.Background())

	// wait for listener to come online
	time.Sleep(100 * time.Millisecond)

	client, err := oapi.NewSnackMgrClient("http://127.0.0.1:24771")
	require.Nil(t, err)

	customerId := uuid.New()
	itemId := uuid.New()
	count := 2

	orderId, err := client.PurchaseOrder(customerId, itemId, count)
	require.Nil(t, err)
	require.NotEqual(t, uuid.Nil, orderId)

}

func newListener() (*http.Server, error) {
	serverAddr := "0.0.0.0:24771"

	backend := backend.NewSnackBackEnd()

	strict := oapi.NewApiHandler(backend)

	r := mux.NewRouter()

	// get an `http.Handler` that we can use
	h := oapi.HandlerFromMuxWithBaseURL(strict, r, "")

	fmt.Printf("starting server and listening on port %s", serverAddr)

	s := &http.Server{
		Handler: h,
		Addr:    serverAddr,
	}

	go func() {
		err := s.ListenAndServe()
		fmt.Printf("\n ended server \n")
		if err != nil {
			fmt.Printf("with error: %s", err.Error())
		}
	}()

	return s, nil
}

func TestPurchaseProcessAcknowledge(t *testing.T) {
	// create server
	// create client

	// order item

	// process item

	// acknowledge item receipt.
}
