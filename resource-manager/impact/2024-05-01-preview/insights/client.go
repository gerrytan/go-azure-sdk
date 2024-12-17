package insights

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightsClient struct {
	Client *resourcemanager.Client
}

func NewInsightsClientWithBaseURI(sdkApi sdkEnv.Api) (*InsightsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "insights", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InsightsClient: %+v", err)
	}

	return &InsightsClient{
		Client: client,
	}, nil
}
