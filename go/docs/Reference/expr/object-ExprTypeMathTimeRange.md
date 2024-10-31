---
title: <span class="badge object-type-struct"></span> ExprTypeMathTimeRange
---
# <span class="badge object-type-struct"></span> ExprTypeMathTimeRange

## Definition

```go
type ExprTypeMathTimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeMathTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeMathTimeRange *ExprTypeMathTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeMathTimeRange` objects.

```go
func (exprTypeMathTimeRange *ExprTypeMathTimeRange) Equals(other ExprTypeMathTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeMathTimeRange` fields for violations and returns them.

```go
func (exprTypeMathTimeRange *ExprTypeMathTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeMathTimeRangeBuilder](./builder-ExprTypeMathTimeRangeBuilder.md)
