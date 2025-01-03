//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config-client.yaml snacks-manager-1.0.0.oapi-3.0.3.yml

package oapi

import (
	"context"
	"fmt"
	"net/http"
)

type SnackMgrClient struct {
	oapi_client *ClientWithResponses
	hc          *http.Client
}

func NewSnackMgrClient(url string) (*SnackMgrClient, error) {

	ret := &SnackMgrClient{
		hc: &http.Client{},
	}

	oapiClient, err := NewClientWithResponses(url, WithHTTPClient(ret.hc))
	if err != nil {
		return nil, err
	}

	ret.oapi_client = oapiClient

	return ret, nil
}

func (c *SnackMgrClient) PurchaseAcknowledge(customerId string, orderId string, transactionNonce string) error {
	received := AllItemsReceived{
		customerId,
		orderId,
		transactionNonce,
	}

	resp, err := c.oapi_client.PostPurchaseAcknowledgeWithResponse(context.TODO(), received, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Errorf("Expected HTTP 200 but received %d", resp.StatusCode())
	}

	// no response to parse, or the response code would have been != 200
	return nil
}

/*
	PostPurchaseAcknowledge(ctx context.Context, body PostPurchaseAcknowledgeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostPurchaseOrderWithBody request with any body
	PostPurchaseOrderWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostPurchaseOrder(ctx context.Context, body PostPurchaseOrderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostPurchaseProcessOrderId request
	PostPurchaseProcessOrderId(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSnacks request
	GetSnacks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSnacksSnackId request
	GetSnacksSnackId(ctx context.Context, snackId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTest request
	GetTest(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
*/
