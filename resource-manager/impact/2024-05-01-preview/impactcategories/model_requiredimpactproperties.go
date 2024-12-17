package impactcategories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredImpactProperties struct {
	AllowedValues *[]string `json:"allowedValues,omitempty"`
	Name          string    `json:"name"`
}
