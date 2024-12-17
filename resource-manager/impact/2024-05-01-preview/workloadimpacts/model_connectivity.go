package workloadimpacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Connectivity struct {
	Port     *int64          `json:"port,omitempty"`
	Protocol *Protocol       `json:"protocol,omitempty"`
	Source   *SourceOrTarget `json:"source,omitempty"`
	Target   *SourceOrTarget `json:"target,omitempty"`
}
