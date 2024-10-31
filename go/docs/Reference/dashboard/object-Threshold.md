---
title: <span class="badge object-type-struct"></span> Threshold
---
# <span class="badge object-type-struct"></span> Threshold

User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded

They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.

## Definition

```go
type Threshold struct {
    // Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
    // Nulls currently appear here when serializing -Infinity to JSON.
    Value *float64 `json:"value"`
    // Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
    Color string `json:"color"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Threshold` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (threshold *Threshold) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Threshold` objects.

```go
func (threshold *Threshold) Equals(other Threshold) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Threshold` fields for violations and returns them.

```go
func (threshold *Threshold) Validate() error
```

