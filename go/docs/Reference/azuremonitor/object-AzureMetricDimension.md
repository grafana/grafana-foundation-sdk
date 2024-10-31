---
title: <span class="badge object-type-struct"></span> AzureMetricDimension
---
# <span class="badge object-type-struct"></span> AzureMetricDimension

## Definition

```go
type AzureMetricDimension struct {
    // Name of Dimension to be filtered on.
    Dimension *string `json:"dimension,omitempty"`
    // String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    Operator *string `json:"operator,omitempty"`
    // Values to match with the filter.
    Filters []string `json:"filters,omitempty"`
    // @deprecated filter is deprecated in favour of filters to support multiselect.
    Filter *string `json:"filter,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMetricDimension` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureMetricDimension *AzureMetricDimension) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureMetricDimension` objects.

```go
func (azureMetricDimension *AzureMetricDimension) Equals(other AzureMetricDimension) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureMetricDimension` fields for violations and returns them.

```go
func (azureMetricDimension *AzureMetricDimension) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureMetricDimensionBuilder](./builder-AzureMetricDimensionBuilder.md)
