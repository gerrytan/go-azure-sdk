// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestParseErrorFromApiResponse_LongRunningOperation(t *testing.T) {
	// LRO specific response
	body := `{"id":"/subscriptions/xxx/providers/Microsoft.AppConfiguration/locations/westeurope/operationsStatus/exSCtZucYyDRejH0Qry3U2xmdvGTx3yogpJ1TObVtSQ","name":"exSCtZucYyDRejH0Qry3U2xmdvGTx3yogpJ1TObVtSQ","status":"Failed","error":{"code":"NameUnavailable","message":"The specified name is already in use.","additionalInfo":[{"type":"ActivityId","info":{"activityId":"0ddb9e62-9386-4040-8f3e-605fb2647d1f"}}]}}`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "0ddb9e62-9386-4040-8f3e-605fb2647d1f",
		Code:         "NameUnavailable",
		FullHttpBody: body,
		Message:      "The specified name is already in use.",
		Status:       "Failed",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_ResourceManagerType1(t *testing.T) {
	// Resource Manager specific response from https://github.com/hashicorp/terraform-provider-azurerm/issues/10858
	body := `{"Code":"Conflict","Message":"Cannot update the site 'func-tfbug-b78d100425' because it uses AlwaysOn feature which is not allowed in the target compute mode.","Target":null,"Details":[{"Message":"Cannot update the site 'func-tfbug-b78d100425' because it uses AlwaysOn feature which is not allowed in the target compute mode."},{"Code":"Conflict"},{"ErrorEntity":{"ExtendedCode":"04068","MessageTemplate":"Cannot update the site '{0}' because it uses AlwaysOn feature which is not allowed in the target compute mode.","Parameters":["func-tfbug-b78d100425"],"Code":"Conflict","Message":"Cannot update the site 'func-tfbug-b78d100425' because it uses AlwaysOn feature which is not allowed in the target compute mode."}}],"Innererror":null}`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "",
		Code:         "Conflict",
		FullHttpBody: body,
		Message:      "Cannot update the site 'func-tfbug-b78d100425' because it uses AlwaysOn feature which is not allowed in the target compute mode.",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_ResourceManagerType2(t *testing.T) {
	// Resource Manager specific response from https://github.com/hashicorp/go-azure-sdk/issues/387
	body := `{"error": {"code": "BadRequest","details": [{"activityId": "abc123","clientRequestId": "xxxx","code": "FailedToStartOperation","message": "Failed to start operation. Verify input and try operation again.","possibleCauses": "Invalid parameters were specified.","recommendedAction": "Verify the input and try again."}],"message": "Failed to start operation. Verify input and try operation again."}}`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "abc123",
		Code:         "FailedToStartOperation",
		FullHttpBody: body,
		Message:      "Failed to start operation. Verify input and try operation again.\nPossible Causes: \"Invalid parameters were specified.\"\nRecommended Action: \"Verify the input and try again.\"",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_EmptyResponseShouldOutputHttpResponseAnyway(t *testing.T) {
	body := ``
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "",
		Code:         "internal-error",
		FullHttpBody: body,
		Message:      "Couldn't parse Azure API Response into a friendly error - please see the original HTTP Response for more details (and file a bug so we can fix this!).",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_UnexpectedHtmlShouldOutputHttpResponseAnyway(t *testing.T) {
	// example from https://github.com/hashicorp/terraform-provider-azurerm/issues/11491
	body := `<!DOCTYPE html><html>    <head>        <title>Runtime Error</title>        <meta name="viewport" content="width=device-width" />        <style>         body {font-family:"Verdana";font-weight:normal;font-size: .7em;color:black;}          p {font-family:"Verdana";font-weight:normal;color:black;margin-top: -5px}         b {font-family:"Verdana";font-weight:bold;color:black;margin-top: -5px}         H1 { font-family:"Verdana";font-weight:normal;font-size:18pt;color:red }         H2 { font-f`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "",
		Code:         "internal-error",
		FullHttpBody: body,
		Message:      "Couldn't parse Azure API Response into a friendly error - please see the original HTTP Response for more details (and file a bug so we can fix this!).",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_UnexpectedLiteralStringShouldOutputHttpResponseAnyway(t *testing.T) {
	// example from https://github.com/hashicorp/terraform-provider-azurerm/issues/1775#issuecomment-927765831
	body := `For first party solution, the solution name has to be convention of solutionType(workspaceName), and the product has to be OMSGallery/solutionType.  For example, product OMSGallery/ChangeTracking, solution name ChangeTracking(myworkspace)`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "",
		Code:         "internal-error",
		FullHttpBody: body,
		Message:      "Couldn't parse Azure API Response into a friendly error - please see the original HTTP Response for more details (and file a bug so we can fix this!).",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_UnexpectedLiteralQuotedStringShouldOutputHttpResponseAnyway(t *testing.T) {
	// example from https://github.com/hashicorp/terraform-provider-azurerm/issues/1775#issuecomment-700385311
	body := `"Solution Not Found : solutionType ADReplication, workspace la-operations-vjhywhmlbmndfroveecsjfkilkkkwswrkoyyggulnwpcaxdni"`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "",
		Code:         "internal-error",
		FullHttpBody: body,
		Message:      "Couldn't parse Azure API Response into a friendly error - please see the original HTTP Response for more details (and file a bug so we can fix this!).",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}

func TestParseErrorFromApiResponse_UnexpectedShouldOutputHttpResponseAnyway(t *testing.T) {
	body := `{"something-went": "wrong"}`
	input := http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}
	expected := Error{
		ActivityId:   "",
		Code:         "internal-error",
		FullHttpBody: body,
		Message:      "Couldn't parse Azure API Response into a friendly error - please see the original HTTP Response for more details (and file a bug so we can fix this!).",
		Status:       "Unknown",
	}
	actual, err := parseErrorFromApiResponse(input)
	if err != nil {
		t.Fatalf("parsing error from api response: %+v", err)
	}
	if actual == nil {
		t.Fatalf("`actual` was nil")
	}
	if !reflect.DeepEqual(*actual, expected) {
		t.Fatalf("expected and actual didn't match. Expected: %+v\n\n Actual: %+v", expected, actual)
	}
}
