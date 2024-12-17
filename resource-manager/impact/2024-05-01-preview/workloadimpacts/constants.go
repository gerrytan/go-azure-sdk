package workloadimpacts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfidenceLevel string

const (
	ConfidenceLevelHigh   ConfidenceLevel = "High"
	ConfidenceLevelLow    ConfidenceLevel = "Low"
	ConfidenceLevelMedium ConfidenceLevel = "Medium"
)

func PossibleValuesForConfidenceLevel() []string {
	return []string{
		string(ConfidenceLevelHigh),
		string(ConfidenceLevelLow),
		string(ConfidenceLevelMedium),
	}
}

func (s *ConfidenceLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfidenceLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfidenceLevel(input string) (*ConfidenceLevel, error) {
	vals := map[string]ConfidenceLevel{
		"high":   ConfidenceLevelHigh,
		"low":    ConfidenceLevelLow,
		"medium": ConfidenceLevelMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfidenceLevel(input)
	return &out, nil
}

type IncidentSource string

const (
	IncidentSourceAzureDevops IncidentSource = "AzureDevops"
	IncidentSourceICM         IncidentSource = "ICM"
	IncidentSourceJira        IncidentSource = "Jira"
	IncidentSourceOther       IncidentSource = "Other"
	IncidentSourceServiceNow  IncidentSource = "ServiceNow"
)

func PossibleValuesForIncidentSource() []string {
	return []string{
		string(IncidentSourceAzureDevops),
		string(IncidentSourceICM),
		string(IncidentSourceJira),
		string(IncidentSourceOther),
		string(IncidentSourceServiceNow),
	}
}

func (s *IncidentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncidentSource(input string) (*IncidentSource, error) {
	vals := map[string]IncidentSource{
		"azuredevops": IncidentSourceAzureDevops,
		"icm":         IncidentSourceICM,
		"jira":        IncidentSourceJira,
		"other":       IncidentSourceOther,
		"servicenow":  IncidentSourceServiceNow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentSource(input)
	return &out, nil
}

type MetricUnit string

const (
	MetricUnitByteSeconds    MetricUnit = "ByteSeconds"
	MetricUnitBytes          MetricUnit = "Bytes"
	MetricUnitBytesPerSecond MetricUnit = "BytesPerSecond"
	MetricUnitCores          MetricUnit = "Cores"
	MetricUnitCount          MetricUnit = "Count"
	MetricUnitCountPerSecond MetricUnit = "CountPerSecond"
	MetricUnitMilliCores     MetricUnit = "MilliCores"
	MetricUnitMilliSeconds   MetricUnit = "MilliSeconds"
	MetricUnitNanoCores      MetricUnit = "NanoCores"
	MetricUnitOther          MetricUnit = "Other"
	MetricUnitPercent        MetricUnit = "Percent"
	MetricUnitSeconds        MetricUnit = "Seconds"
)

func PossibleValuesForMetricUnit() []string {
	return []string{
		string(MetricUnitByteSeconds),
		string(MetricUnitBytes),
		string(MetricUnitBytesPerSecond),
		string(MetricUnitCores),
		string(MetricUnitCount),
		string(MetricUnitCountPerSecond),
		string(MetricUnitMilliCores),
		string(MetricUnitMilliSeconds),
		string(MetricUnitNanoCores),
		string(MetricUnitOther),
		string(MetricUnitPercent),
		string(MetricUnitSeconds),
	}
}

func (s *MetricUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricUnit(input string) (*MetricUnit, error) {
	vals := map[string]MetricUnit{
		"byteseconds":    MetricUnitByteSeconds,
		"bytes":          MetricUnitBytes,
		"bytespersecond": MetricUnitBytesPerSecond,
		"cores":          MetricUnitCores,
		"count":          MetricUnitCount,
		"countpersecond": MetricUnitCountPerSecond,
		"millicores":     MetricUnitMilliCores,
		"milliseconds":   MetricUnitMilliSeconds,
		"nanocores":      MetricUnitNanoCores,
		"other":          MetricUnitOther,
		"percent":        MetricUnitPercent,
		"seconds":        MetricUnitSeconds,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricUnit(input)
	return &out, nil
}

type Protocol string

const (
	ProtocolFTP   Protocol = "FTP"
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolHTTPS Protocol = "HTTPS"
	ProtocolOther Protocol = "Other"
	ProtocolRDP   Protocol = "RDP"
	ProtocolSSH   Protocol = "SSH"
	ProtocolTCP   Protocol = "TCP"
	ProtocolUDP   Protocol = "UDP"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolFTP),
		string(ProtocolHTTP),
		string(ProtocolHTTPS),
		string(ProtocolOther),
		string(ProtocolRDP),
		string(ProtocolSSH),
		string(ProtocolTCP),
		string(ProtocolUDP),
	}
}

func (s *Protocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"ftp":   ProtocolFTP,
		"http":  ProtocolHTTP,
		"https": ProtocolHTTPS,
		"other": ProtocolOther,
		"rdp":   ProtocolRDP,
		"ssh":   ProtocolSSH,
		"tcp":   ProtocolTCP,
		"udp":   ProtocolUDP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type Toolset string

const (
	ToolsetARM       Toolset = "ARM"
	ToolsetAnsible   Toolset = "Ansible"
	ToolsetChef      Toolset = "Chef"
	ToolsetOther     Toolset = "Other"
	ToolsetPortal    Toolset = "Portal"
	ToolsetPuppet    Toolset = "Puppet"
	ToolsetSDK       Toolset = "SDK"
	ToolsetShell     Toolset = "Shell"
	ToolsetTerraform Toolset = "Terraform"
)

func PossibleValuesForToolset() []string {
	return []string{
		string(ToolsetARM),
		string(ToolsetAnsible),
		string(ToolsetChef),
		string(ToolsetOther),
		string(ToolsetPortal),
		string(ToolsetPuppet),
		string(ToolsetSDK),
		string(ToolsetShell),
		string(ToolsetTerraform),
	}
}

func (s *Toolset) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseToolset(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseToolset(input string) (*Toolset, error) {
	vals := map[string]Toolset{
		"arm":       ToolsetARM,
		"ansible":   ToolsetAnsible,
		"chef":      ToolsetChef,
		"other":     ToolsetOther,
		"portal":    ToolsetPortal,
		"puppet":    ToolsetPuppet,
		"sdk":       ToolsetSDK,
		"shell":     ToolsetShell,
		"terraform": ToolsetTerraform,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Toolset(input)
	return &out, nil
}
