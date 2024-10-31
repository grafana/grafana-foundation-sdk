---
title: <span class="badge object-type-struct"></span> AzureTracesFilter
---
# <span class="badge object-type-struct"></span> AzureTracesFilter

## Definition

```go
type AzureTracesFilter struct {
    // Property name, auto-populated based on available traces.
    Property string `json:"property"`
    // Comparison operator to use. Either equals or not equals.
    Operation string `json:"operation"`
    // Values to filter by.
    Filters []string `json:"filters"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureTracesFilter` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureTracesFilter *AzureTracesFilter) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureTracesFilter` objects.

```go
func (azureTracesFilter *AzureTracesFilter) Equals(other AzureTracesFilter) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureTracesFilter` fields for violations and returns them.

```go
func (azureTracesFilter *AzureTracesFilter) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureTracesFilterBuilder](./builder-AzureTracesFilterBuilder.md)
