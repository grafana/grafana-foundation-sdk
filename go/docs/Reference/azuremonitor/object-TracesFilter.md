---
title: <span class="badge object-type-struct"></span> TracesFilter
---
# <span class="badge object-type-struct"></span> TracesFilter

## Definition

```go
type TracesFilter struct {
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

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TracesFilter` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tracesFilter *TracesFilter) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TracesFilter` objects.

```go
func (tracesFilter *TracesFilter) Equals(other TracesFilter) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TracesFilter` fields for violations and returns them.

```go
func (tracesFilter *TracesFilter) Validate() error
```

## See also

 * <span class="badge builder"></span> [TracesFilterBuilder](./builder-TracesFilterBuilder.md)
