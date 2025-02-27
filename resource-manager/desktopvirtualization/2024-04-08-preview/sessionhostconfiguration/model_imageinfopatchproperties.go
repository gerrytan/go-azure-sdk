package sessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageInfoPatchProperties struct {
	CustomInfo      *CustomInfoPatchProperties      `json:"customInfo,omitempty"`
	MarketplaceInfo *MarketplaceInfoPatchProperties `json:"marketplaceInfo,omitempty"`
	Type            *Type                           `json:"type,omitempty"`
}
