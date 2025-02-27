package loadbalancers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LoadBalancingRuleId{}

func TestNewLoadBalancingRuleID(t *testing.T) {
	id := NewLoadBalancingRuleID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "loadBalancerName", "loadBalancingRuleName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "resourceGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "resourceGroupName")
	}

	if id.LoadBalancerName != "loadBalancerName" {
		t.Fatalf("Expected %q but got %q for Segment 'LoadBalancerName'", id.LoadBalancerName, "loadBalancerName")
	}

	if id.LoadBalancingRuleName != "loadBalancingRuleName" {
		t.Fatalf("Expected %q but got %q for Segment 'LoadBalancingRuleName'", id.LoadBalancingRuleName, "loadBalancingRuleName")
	}
}

func TestFormatLoadBalancingRuleID(t *testing.T) {
	actual := NewLoadBalancingRuleID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "loadBalancerName", "loadBalancingRuleName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules/loadBalancingRuleName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLoadBalancingRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LoadBalancingRuleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules/loadBalancingRuleName",
			Expected: &LoadBalancingRuleId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "resourceGroupName",
				LoadBalancerName:      "loadBalancerName",
				LoadBalancingRuleName: "loadBalancingRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules/loadBalancingRuleName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLoadBalancingRuleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.LoadBalancerName != v.Expected.LoadBalancerName {
			t.Fatalf("Expected %q but got %q for LoadBalancerName", v.Expected.LoadBalancerName, actual.LoadBalancerName)
		}

		if actual.LoadBalancingRuleName != v.Expected.LoadBalancingRuleName {
			t.Fatalf("Expected %q but got %q for LoadBalancingRuleName", v.Expected.LoadBalancingRuleName, actual.LoadBalancingRuleName)
		}

	}
}

func TestParseLoadBalancingRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LoadBalancingRuleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.nEtWoRk/lOaDbAlAnCeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.nEtWoRk/lOaDbAlAnCeRs/lOaDbAlAnCeRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.nEtWoRk/lOaDbAlAnCeRs/lOaDbAlAnCeRnAmE/lOaDbAlAnCiNgRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules/loadBalancingRuleName",
			Expected: &LoadBalancingRuleId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "resourceGroupName",
				LoadBalancerName:      "loadBalancerName",
				LoadBalancingRuleName: "loadBalancingRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.Network/loadBalancers/loadBalancerName/loadBalancingRules/loadBalancingRuleName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.nEtWoRk/lOaDbAlAnCeRs/lOaDbAlAnCeRnAmE/lOaDbAlAnCiNgRuLeS/lOaDbAlAnCiNgRuLeNaMe",
			Expected: &LoadBalancingRuleId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "rEsOuRcEgRoUpNaMe",
				LoadBalancerName:      "lOaDbAlAnCeRnAmE",
				LoadBalancingRuleName: "lOaDbAlAnCiNgRuLeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.nEtWoRk/lOaDbAlAnCeRs/lOaDbAlAnCeRnAmE/lOaDbAlAnCiNgRuLeS/lOaDbAlAnCiNgRuLeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLoadBalancingRuleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.LoadBalancerName != v.Expected.LoadBalancerName {
			t.Fatalf("Expected %q but got %q for LoadBalancerName", v.Expected.LoadBalancerName, actual.LoadBalancerName)
		}

		if actual.LoadBalancingRuleName != v.Expected.LoadBalancingRuleName {
			t.Fatalf("Expected %q but got %q for LoadBalancingRuleName", v.Expected.LoadBalancingRuleName, actual.LoadBalancingRuleName)
		}

	}
}

func TestSegmentsForLoadBalancingRuleId(t *testing.T) {
	segments := LoadBalancingRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LoadBalancingRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
