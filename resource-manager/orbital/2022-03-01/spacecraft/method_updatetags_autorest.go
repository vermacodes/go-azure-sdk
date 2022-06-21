package spacecraft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTagsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateTags ...
func (c SpacecraftClient) UpdateTags(ctx context.Context, id SpacecraftId, input TagsObject) (result UpdateTagsOperationResponse, err error) {
	req, err := c.preparerForUpdateTags(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spacecraft.SpacecraftClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateTags(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spacecraft.SpacecraftClient", "UpdateTags", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateTagsThenPoll performs UpdateTags then polls until it's completed
func (c SpacecraftClient) UpdateTagsThenPoll(ctx context.Context, id SpacecraftId, input TagsObject) error {
	result, err := c.UpdateTags(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateTags: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateTags: %+v", err)
	}

	return nil
}

// preparerForUpdateTags prepares the UpdateTags request.
func (c SpacecraftClient) preparerForUpdateTags(ctx context.Context, id SpacecraftId, input TagsObject) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateTags sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (c SpacecraftClient) senderForUpdateTags(ctx context.Context, req *http.Request) (future UpdateTagsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
