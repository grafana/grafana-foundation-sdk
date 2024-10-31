---
title: <span class="badge object-type-struct"></span> ExprTypeReduceTimeRange
---
# <span class="badge object-type-struct"></span> ExprTypeReduceTimeRange

## Definition

```go
type ExprTypeReduceTimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeReduceTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeReduceTimeRange *ExprTypeReduceTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeReduceTimeRange` objects.

```go
func (exprTypeReduceTimeRange *ExprTypeReduceTimeRange) Equals(other ExprTypeReduceTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeReduceTimeRange` fields for violations and returns them.

```go
func (exprTypeReduceTimeRange *ExprTypeReduceTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeReduceTimeRangeBuilder](./builder-ExprTypeReduceTimeRangeBuilder.md)
