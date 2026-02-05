---
title: <span class="badge object-type-struct"></span> ExprTypeSqlTimeRange
---
# <span class="badge object-type-struct"></span> ExprTypeSqlTimeRange

## Definition

```go
type ExprTypeSqlTimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeSqlTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeSqlTimeRange *ExprTypeSqlTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeSqlTimeRange` objects.

```go
func (exprTypeSqlTimeRange *ExprTypeSqlTimeRange) Equals(other ExprTypeSqlTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeSqlTimeRange` fields for violations and returns them.

```go
func (exprTypeSqlTimeRange *ExprTypeSqlTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeSqlTimeRangeBuilder](./builder-ExprTypeSqlTimeRangeBuilder.md)
