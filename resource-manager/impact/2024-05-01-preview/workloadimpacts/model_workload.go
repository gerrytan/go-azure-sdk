package workloadimpacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Workload struct {
	Context *string  `json:"context,omitempty"`
	Toolset *Toolset `json:"toolset,omitempty"`
}
