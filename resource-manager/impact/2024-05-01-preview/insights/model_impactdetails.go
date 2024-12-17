package insights

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactDetails struct {
	EndTime            *string `json:"endTime,omitempty"`
	ImpactId           string  `json:"impactId"`
	ImpactedResourceId string  `json:"impactedResourceId"`
	StartTime          string  `json:"startTime"`
}

func (o *ImpactDetails) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ImpactDetails) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *ImpactDetails) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ImpactDetails) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
