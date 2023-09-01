package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Volume
}

type VolumesListByResourceGroupCompleteResult struct {
	Items []Volume
}

// VolumesListByResourceGroup ...
func (c NetworkcloudsClient) VolumesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result VolumesListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/volumes", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Volume `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VolumesListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) VolumesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (VolumesListByResourceGroupCompleteResult, error) {
	return c.VolumesListByResourceGroupCompleteMatchingPredicate(ctx, id, VolumeOperationPredicate{})
}

// VolumesListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) VolumesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate VolumeOperationPredicate) (result VolumesListByResourceGroupCompleteResult, err error) {
	items := make([]Volume, 0)

	resp, err := c.VolumesListByResourceGroup(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = VolumesListByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
