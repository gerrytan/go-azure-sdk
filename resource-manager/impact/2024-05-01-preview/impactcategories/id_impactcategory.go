package impactcategories

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ImpactCategoryId{})
}

var _ resourceids.ResourceId = &ImpactCategoryId{}

// ImpactCategoryId is a struct representing the Resource ID for a Impact Category
type ImpactCategoryId struct {
	SubscriptionId     string
	ImpactCategoryName string
}

// NewImpactCategoryID returns a new ImpactCategoryId struct
func NewImpactCategoryID(subscriptionId string, impactCategoryName string) ImpactCategoryId {
	return ImpactCategoryId{
		SubscriptionId:     subscriptionId,
		ImpactCategoryName: impactCategoryName,
	}
}

// ParseImpactCategoryID parses 'input' into a ImpactCategoryId
func ParseImpactCategoryID(input string) (*ImpactCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ImpactCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ImpactCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseImpactCategoryIDInsensitively parses 'input' case-insensitively into a ImpactCategoryId
// note: this method should only be used for API response data and not user input
func ParseImpactCategoryIDInsensitively(input string) (*ImpactCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ImpactCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ImpactCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ImpactCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ImpactCategoryName, ok = input.Parsed["impactCategoryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "impactCategoryName", input)
	}

	return nil
}

// ValidateImpactCategoryID checks that 'input' can be parsed as a Impact Category ID
func ValidateImpactCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseImpactCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Impact Category ID
func (id ImpactCategoryId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Impact/impactCategories/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ImpactCategoryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Impact Category ID
func (id ImpactCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftImpact", "Microsoft.Impact", "Microsoft.Impact"),
		resourceids.StaticSegment("staticImpactCategories", "impactCategories", "impactCategories"),
		resourceids.UserSpecifiedSegment("impactCategoryName", "impactCategoryName"),
	}
}

// String returns a human-readable description of this Impact Category ID
func (id ImpactCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Impact Category Name: %q", id.ImpactCategoryName),
	}
	return fmt.Sprintf("Impact Category (%s)", strings.Join(components, "\n"))
}
