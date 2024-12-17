package v2024_05_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/connectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/impactcategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/insights"
	"github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/workloadimpacts"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Connectors       *connectors.ConnectorsClient
	ImpactCategories *impactcategories.ImpactCategoriesClient
	Insights         *insights.InsightsClient
	WorkloadImpacts  *workloadimpacts.WorkloadImpactsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	connectorsClient, err := connectors.NewConnectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Connectors client: %+v", err)
	}
	configureFunc(connectorsClient.Client)

	impactCategoriesClient, err := impactcategories.NewImpactCategoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ImpactCategories client: %+v", err)
	}
	configureFunc(impactCategoriesClient.Client)

	insightsClient, err := insights.NewInsightsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Insights client: %+v", err)
	}
	configureFunc(insightsClient.Client)

	workloadImpactsClient, err := workloadimpacts.NewWorkloadImpactsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadImpacts client: %+v", err)
	}
	configureFunc(workloadImpactsClient.Client)

	return &Client{
		Connectors:       connectorsClient,
		ImpactCategories: impactCategoriesClient,
		Insights:         insightsClient,
		WorkloadImpacts:  workloadImpactsClient,
	}, nil
}
