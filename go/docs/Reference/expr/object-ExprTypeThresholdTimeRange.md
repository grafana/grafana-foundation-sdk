---
title: <span class="badge object-type-struct"></span> ExprTypeThresholdTimeRange
---
# <span class="badge object-type-struct"></span> ExprTypeThresholdTimeRange

## Definition

```go
type ExprTypeThresholdTimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeThresholdTimeRange *ExprTypeThresholdTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeThresholdTimeRange` objects.

```go
func (exprTypeThresholdTimeRange *ExprTypeThresholdTimeRange) Equals(other ExprTypeThresholdTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeThresholdTimeRange` fields for violations and returns them.

```go
func (exprTypeThresholdTimeRange *ExprTypeThresholdTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeThresholdTimeRangeBuilder](./builder-ExprTypeThresholdTimeRangeBuilder.md)
