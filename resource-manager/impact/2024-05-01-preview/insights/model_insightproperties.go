package insights

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightProperties struct {
	AdditionalDetails *interface{}       `json:"additionalDetails,omitempty"`
	Category          string             `json:"category"`
	Content           Content            `json:"content"`
	EventId           *string            `json:"eventId,omitempty"`
	EventTime         *string            `json:"eventTime,omitempty"`
	GroupId           *string            `json:"groupId,omitempty"`
	Impact            ImpactDetails      `json:"impact"`
	InsightUniqueId   string             `json:"insightUniqueId"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Status            *string            `json:"status,omitempty"`
}

func (o *InsightProperties) GetEventTimeAsTime() (*time.Time, error) {
	if o.EventTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventTime, "2006-01-02T15:04:05Z07:00")
}

func (o *InsightProperties) SetEventTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EventTime = &formatted
}
