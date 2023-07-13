package applyupdate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyUpdateClient struct {
	Client *resourcemanager.Client
}

func NewApplyUpdateClientWithBaseURI(api environments.Api) (*ApplyUpdateClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "applyupdate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplyUpdateClient: %+v", err)
	}

	return &ApplyUpdateClient{
		Client: client,
	}, nil
}
