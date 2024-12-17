package impactcategories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactCategoryProperties struct {
	CategoryId               string                      `json:"categoryId"`
	Description              *string                     `json:"description,omitempty"`
	ParentCategoryId         *string                     `json:"parentCategoryId,omitempty"`
	ProvisioningState        *ProvisioningState          `json:"provisioningState,omitempty"`
	RequiredImpactProperties *[]RequiredImpactProperties `json:"requiredImpactProperties,omitempty"`
}
