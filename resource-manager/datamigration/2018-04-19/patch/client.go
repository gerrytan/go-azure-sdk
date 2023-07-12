package patch

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PATCHClient struct {
	Client *resourcemanager.Client
}

func NewPATCHClientWithBaseURI(api environments.Api) (*PATCHClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "patch", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PATCHClient: %+v", err)
	}

	return &PATCHClient{
		Client: client,
	}, nil
}
