---
title: <span class="badge object-type-struct"></span> AzureMonitorResource
---
# <span class="badge object-type-struct"></span> AzureMonitorResource

## Definition

```go
type AzureMonitorResource struct {
    Subscription *string `json:"subscription,omitempty"`
    ResourceGroup *string `json:"resourceGroup,omitempty"`
    ResourceName *string `json:"resourceName,omitempty"`
    MetricNamespace *string `json:"metricNamespace,omitempty"`
    Region *string `json:"region,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMonitorResource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureMonitorResource *AzureMonitorResource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureMonitorResource` objects.

```go
func (azureMonitorResource *AzureMonitorResource) Equals(other AzureMonitorResource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureMonitorResource` fields for violations and returns them.

```go
func (azureMonitorResource *AzureMonitorResource) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureMonitorResourceBuilder](./builder-AzureMonitorResourceBuilder.md)
