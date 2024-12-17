package impactcategories

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactCategoriesClient struct {
	Client *resourcemanager.Client
}

func NewImpactCategoriesClientWithBaseURI(sdkApi sdkEnv.Api) (*ImpactCategoriesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "impactcategories", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ImpactCategoriesClient: %+v", err)
	}

	return &ImpactCategoriesClient{
		Client: client,
	}, nil
}
