package netappresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceClient struct {
	Client *resourcemanager.Client
}

func NewNetAppResourceClientWithBaseURI(api environments.Api) (*NetAppResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "netappresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetAppResourceClient: %+v", err)
	}

	return &NetAppResourceClient{
		Client: client,
	}, nil
}
