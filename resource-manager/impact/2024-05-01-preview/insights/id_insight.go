package insights

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&InsightId{})
}

var _ resourceids.ResourceId = &InsightId{}

// InsightId is a struct representing the Resource ID for a Insight
type InsightId struct {
	SubscriptionId     string
	WorkloadImpactName string
	InsightName        string
}

// NewInsightID returns a new InsightId struct
func NewInsightID(subscriptionId string, workloadImpactName string, insightName string) InsightId {
	return InsightId{
		SubscriptionId:     subscriptionId,
		WorkloadImpactName: workloadImpactName,
		InsightName:        insightName,
	}
}

// ParseInsightID parses 'input' into a InsightId
func ParseInsightID(input string) (*InsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInsightIDInsensitively parses 'input' case-insensitively into a InsightId
// note: this method should only be used for API response data and not user input
func ParseInsightIDInsensitively(input string) (*InsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.WorkloadImpactName, ok = input.Parsed["workloadImpactName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workloadImpactName", input)
	}

	if id.InsightName, ok = input.Parsed["insightName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "insightName", input)
	}

	return nil
}

// ValidateInsightID checks that 'input' can be parsed as a Insight ID
func ValidateInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Insight ID
func (id InsightId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Impact/workloadImpacts/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.WorkloadImpactName, id.InsightName)
}

// Segments returns a slice of Resource ID Segments which comprise this Insight ID
func (id InsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftImpact", "Microsoft.Impact", "Microsoft.Impact"),
		resourceids.StaticSegment("staticWorkloadImpacts", "workloadImpacts", "workloadImpacts"),
		resourceids.UserSpecifiedSegment("workloadImpactName", "workloadImpactName"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("insightName", "insightName"),
	}
}

// String returns a human-readable description of this Insight ID
func (id InsightId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Workload Impact Name: %q", id.WorkloadImpactName),
		fmt.Sprintf("Insight Name: %q", id.InsightName),
	}
	return fmt.Sprintf("Insight (%s)", strings.Join(components, "\n"))
}
