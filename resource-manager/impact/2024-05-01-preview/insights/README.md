
## `github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/insights` Documentation

The `insights` SDK allows for interaction with Azure Resource Manager `impact` (API Version `2024-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/insights"
```


### Client Initialization

```go
client := insights.NewInsightsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InsightsClient.Create`

```go
ctx := context.TODO()
id := insights.NewInsightID("12345678-1234-9876-4563-123456789012", "workloadImpactName", "insightName")

payload := insights.Insight{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InsightsClient.Delete`

```go
ctx := context.TODO()
id := insights.NewInsightID("12345678-1234-9876-4563-123456789012", "workloadImpactName", "insightName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InsightsClient.Get`

```go
ctx := context.TODO()
id := insights.NewInsightID("12345678-1234-9876-4563-123456789012", "workloadImpactName", "insightName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InsightsClient.ListBySubscription`

```go
ctx := context.TODO()
id := insights.NewWorkloadImpactID("12345678-1234-9876-4563-123456789012", "workloadImpactName")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
