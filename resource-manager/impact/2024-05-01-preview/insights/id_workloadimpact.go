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
	recaser.RegisterResourceId(&WorkloadImpactId{})
}

var _ resourceids.ResourceId = &WorkloadImpactId{}

// WorkloadImpactId is a struct representing the Resource ID for a Workload Impact
type WorkloadImpactId struct {
	SubscriptionId     string
	WorkloadImpactName string
}

// NewWorkloadImpactID returns a new WorkloadImpactId struct
func NewWorkloadImpactID(subscriptionId string, workloadImpactName string) WorkloadImpactId {
	return WorkloadImpactId{
		SubscriptionId:     subscriptionId,
		WorkloadImpactName: workloadImpactName,
	}
}

// ParseWorkloadImpactID parses 'input' into a WorkloadImpactId
func ParseWorkloadImpactID(input string) (*WorkloadImpactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkloadImpactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkloadImpactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkloadImpactIDInsensitively parses 'input' case-insensitively into a WorkloadImpactId
// note: this method should only be used for API response data and not user input
func ParseWorkloadImpactIDInsensitively(input string) (*WorkloadImpactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkloadImpactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkloadImpactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkloadImpactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.WorkloadImpactName, ok = input.Parsed["workloadImpactName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workloadImpactName", input)
	}

	return nil
}

// ValidateWorkloadImpactID checks that 'input' can be parsed as a Workload Impact ID
func ValidateWorkloadImpactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkloadImpactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workload Impact ID
func (id WorkloadImpactId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Impact/workloadImpacts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.WorkloadImpactName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workload Impact ID
func (id WorkloadImpactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftImpact", "Microsoft.Impact", "Microsoft.Impact"),
		resourceids.StaticSegment("staticWorkloadImpacts", "workloadImpacts", "workloadImpacts"),
		resourceids.UserSpecifiedSegment("workloadImpactName", "workloadImpactName"),
	}
}

// String returns a human-readable description of this Workload Impact ID
func (id WorkloadImpactId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Workload Impact Name: %q", id.WorkloadImpactName),
	}
	return fmt.Sprintf("Workload Impact (%s)", strings.Join(components, "\n"))
}
