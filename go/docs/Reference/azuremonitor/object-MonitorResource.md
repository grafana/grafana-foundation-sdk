---
title: <span class="badge object-type-struct"></span> MonitorResource
---
# <span class="badge object-type-struct"></span> MonitorResource

## Definition

```go
type MonitorResource struct {
    Subscription *string `json:"subscription,omitempty"`
    ResourceGroup *string `json:"resourceGroup,omitempty"`
    ResourceName *string `json:"resourceName,omitempty"`
    MetricNamespace *string `json:"metricNamespace,omitempty"`
    Region *string `json:"region,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MonitorResource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (monitorResource *MonitorResource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MonitorResource` objects.

```go
func (monitorResource *MonitorResource) Equals(other MonitorResource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MonitorResource` fields for violations and returns them.

```go
func (monitorResource *MonitorResource) Validate() error
```

## See also

 * <span class="badge builder"></span> [MonitorResourceBuilder](./builder-MonitorResourceBuilder.md)
