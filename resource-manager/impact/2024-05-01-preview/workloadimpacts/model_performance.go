package workloadimpacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Performance struct {
	Actual             *float64            `json:"actual,omitempty"`
	Expected           *float64            `json:"expected,omitempty"`
	ExpectedValueRange *ExpectedValueRange `json:"expectedValueRange,omitempty"`
	MetricName         *string             `json:"metricName,omitempty"`
	Unit               *MetricUnit         `json:"unit,omitempty"`
}
