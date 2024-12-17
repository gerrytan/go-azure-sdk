
## `github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/workloadimpacts` Documentation

The `workloadimpacts` SDK allows for interaction with Azure Resource Manager `impact` (API Version `2024-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/impact/2024-05-01-preview/workloadimpacts"
```


### Client Initialization

```go
client := workloadimpacts.NewWorkloadImpactsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadImpactsClient.Create`

```go
ctx := context.TODO()
id := workloadimpacts.NewWorkloadImpactID("12345678-1234-9876-4563-123456789012", "workloadImpactName")

payload := workloadimpacts.WorkloadImpact{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadImpactsClient.Delete`

```go
ctx := context.TODO()
id := workloadimpacts.NewWorkloadImpactID("12345678-1234-9876-4563-123456789012", "workloadImpactName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadImpactsClient.Get`

```go
ctx := context.TODO()
id := workloadimpacts.NewWorkloadImpactID("12345678-1234-9876-4563-123456789012", "workloadImpactName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadImpactsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
