package inferencedeltamodel

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeltaModelCurrentState struct {
	Count            *int64  `json:"count,omitempty"`
	SampleInstanceID *string `json:"sampleInstanceID,omitempty"`
	Status           *string `json:"status,omitempty"`
}
