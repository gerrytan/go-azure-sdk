package workloadimpacts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadImpactProperties struct {
	AdditionalProperties  *interface{}           `json:"additionalProperties,omitempty"`
	ArmCorrelationIds     *[]string              `json:"armCorrelationIds,omitempty"`
	ClientIncidentDetails *ClientIncidentDetails `json:"clientIncidentDetails,omitempty"`
	ConfidenceLevel       *ConfidenceLevel       `json:"confidenceLevel,omitempty"`
	Connectivity          *Connectivity          `json:"connectivity,omitempty"`
	EndDateTime           *string                `json:"endDateTime,omitempty"`
	ErrorDetails          *ErrorDetailProperties `json:"errorDetails,omitempty"`
	ImpactCategory        string                 `json:"impactCategory"`
	ImpactDescription     *string                `json:"impactDescription,omitempty"`
	ImpactGroupId         *string                `json:"impactGroupId,omitempty"`
	ImpactUniqueId        *string                `json:"impactUniqueId,omitempty"`
	ImpactedResourceId    string                 `json:"impactedResourceId"`
	Performance           *[]Performance         `json:"performance,omitempty"`
	ProvisioningState     *ProvisioningState     `json:"provisioningState,omitempty"`
	ReportedTimeUtc       *string                `json:"reportedTimeUtc,omitempty"`
	StartDateTime         string                 `json:"startDateTime"`
	Workload              *Workload              `json:"workload,omitempty"`
}

func (o *WorkloadImpactProperties) GetEndDateTimeAsTime() (*time.Time, error) {
	if o.EndDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkloadImpactProperties) SetEndDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDateTime = &formatted
}

func (o *WorkloadImpactProperties) GetReportedTimeUtcAsTime() (*time.Time, error) {
	if o.ReportedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReportedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkloadImpactProperties) SetReportedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReportedTimeUtc = &formatted
}

func (o *WorkloadImpactProperties) GetStartDateTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkloadImpactProperties) SetStartDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDateTime = formatted
}
