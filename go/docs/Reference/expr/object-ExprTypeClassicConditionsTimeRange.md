---
title: <span class="badge object-type-struct"></span> ExprTypeClassicConditionsTimeRange
---
# <span class="badge object-type-struct"></span> ExprTypeClassicConditionsTimeRange

## Definition

```go
type ExprTypeClassicConditionsTimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeClassicConditionsTimeRange *ExprTypeClassicConditionsTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeClassicConditionsTimeRange` objects.

```go
func (exprTypeClassicConditionsTimeRange *ExprTypeClassicConditionsTimeRange) Equals(other ExprTypeClassicConditionsTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsTimeRange` fields for violations and returns them.

```go
func (exprTypeClassicConditionsTimeRange *ExprTypeClassicConditionsTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeClassicConditionsTimeRangeBuilder](./builder-ExprTypeClassicConditionsTimeRangeBuilder.md)
