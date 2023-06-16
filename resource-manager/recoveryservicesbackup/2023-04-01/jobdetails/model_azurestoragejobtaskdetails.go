package jobdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureStorageJobTaskDetails struct {
	Status *string `json:"status,omitempty"`
	TaskId *string `json:"taskId,omitempty"`
}
