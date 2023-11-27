package workflowversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VersionId{}

// VersionId is a struct representing the Resource ID for a Version
type VersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkflowName      string
	VersionId         string
}

// NewVersionID returns a new VersionId struct
func NewVersionID(subscriptionId string, resourceGroupName string, workflowName string, versionId string) VersionId {
	return VersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkflowName:      workflowName,
		VersionId:         versionId,
	}
}

// ParseVersionID parses 'input' into a VersionId
func ParseVersionID(input string) (*VersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVersionIDInsensitively parses 'input' case-insensitively into a VersionId
// note: this method should only be used for API response data and not user input
func ParseVersionIDInsensitively(input string) (*VersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkflowName, ok = input.Parsed["workflowName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowName", input)
	}

	if id.VersionId, ok = input.Parsed["versionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionId", input)
	}

	return nil
}

// ValidateVersionID checks that 'input' can be parsed as a Version ID
func ValidateVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Version ID
func (id VersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Logic/workflows/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkflowName, id.VersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Version ID
func (id VersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftLogic", "Microsoft.Logic", "Microsoft.Logic"),
		resourceids.StaticSegment("staticWorkflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowName", "workflowValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionId", "versionIdValue"),
	}
}

// String returns a human-readable description of this Version ID
func (id VersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workflow Name: %q", id.WorkflowName),
		fmt.Sprintf("Version: %q", id.VersionId),
	}
	return fmt.Sprintf("Version (%s)", strings.Join(components, "\n"))
}
