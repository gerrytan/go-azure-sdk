package insights

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InsightId{}

func TestNewInsightID(t *testing.T) {
	id := NewInsightID("12345678-1234-9876-4563-123456789012", "workloadImpactName", "insightName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.WorkloadImpactName != "workloadImpactName" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkloadImpactName'", id.WorkloadImpactName, "workloadImpactName")
	}

	if id.InsightName != "insightName" {
		t.Fatalf("Expected %q but got %q for Segment 'InsightName'", id.InsightName, "insightName")
	}
}

func TestFormatInsightID(t *testing.T) {
	actual := NewInsightID("12345678-1234-9876-4563-123456789012", "workloadImpactName", "insightName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights/insightName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InsightId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights/insightName",
			Expected: &InsightId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				WorkloadImpactName: "workloadImpactName",
				InsightName:        "insightName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights/insightName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInsightID(v.Input)
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

		if actual.WorkloadImpactName != v.Expected.WorkloadImpactName {
			t.Fatalf("Expected %q but got %q for WorkloadImpactName", v.Expected.WorkloadImpactName, actual.WorkloadImpactName)
		}

		if actual.InsightName != v.Expected.InsightName {
			t.Fatalf("Expected %q but got %q for InsightName", v.Expected.InsightName, actual.InsightName)
		}

	}
}

func TestParseInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InsightId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.iMpAcT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.iMpAcT/wOrKlOaDiMpAcTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.iMpAcT/wOrKlOaDiMpAcTs/wOrKlOaDiMpAcTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.iMpAcT/wOrKlOaDiMpAcTs/wOrKlOaDiMpAcTnAmE/iNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights/insightName",
			Expected: &InsightId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				WorkloadImpactName: "workloadImpactName",
				InsightName:        "insightName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Impact/workloadImpacts/workloadImpactName/insights/insightName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.iMpAcT/wOrKlOaDiMpAcTs/wOrKlOaDiMpAcTnAmE/iNsIgHtS/iNsIgHtNaMe",
			Expected: &InsightId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				WorkloadImpactName: "wOrKlOaDiMpAcTnAmE",
				InsightName:        "iNsIgHtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.iMpAcT/wOrKlOaDiMpAcTs/wOrKlOaDiMpAcTnAmE/iNsIgHtS/iNsIgHtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInsightIDInsensitively(v.Input)
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

		if actual.WorkloadImpactName != v.Expected.WorkloadImpactName {
			t.Fatalf("Expected %q but got %q for WorkloadImpactName", v.Expected.WorkloadImpactName, actual.WorkloadImpactName)
		}

		if actual.InsightName != v.Expected.InsightName {
			t.Fatalf("Expected %q but got %q for InsightName", v.Expected.InsightName, actual.InsightName)
		}

	}
}

func TestSegmentsForInsightId(t *testing.T) {
	segments := InsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("InsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
