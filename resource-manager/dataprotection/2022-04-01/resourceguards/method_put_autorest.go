package resourceguards

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type PutOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceGuardResource
}

// Put ...
func (c ResourceGuardsClient) Put(ctx context.Context, id ResourceGuardId, input ResourceGuardResource) (result PutOperationResponse, err error) {
	req, err := c.preparerForPut(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceguards.ResourceGuardsClient", "Put", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceguards.ResourceGuardsClient", "Put", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPut(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceguards.ResourceGuardsClient", "Put", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPut prepares the Put request.
func (c ResourceGuardsClient) preparerForPut(ctx context.Context, id ResourceGuardId, input ResourceGuardResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPut handles the response to the Put request. The method always
// closes the http.Response Body.
func (c ResourceGuardsClient) responderForPut(resp *http.Response) (result PutOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
