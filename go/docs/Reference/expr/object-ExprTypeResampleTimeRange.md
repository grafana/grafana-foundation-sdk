---
title: <span class="badge object-type-struct"></span> ExprTypeResampleTimeRange
---
# <span class="badge object-type-struct"></span> ExprTypeResampleTimeRange

## Definition

```go
type ExprTypeResampleTimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeResampleTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeResampleTimeRange *ExprTypeResampleTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeResampleTimeRange` objects.

```go
func (exprTypeResampleTimeRange *ExprTypeResampleTimeRange) Equals(other ExprTypeResampleTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeResampleTimeRange` fields for violations and returns them.

```go
func (exprTypeResampleTimeRange *ExprTypeResampleTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeResampleTimeRangeBuilder](./builder-ExprTypeResampleTimeRangeBuilder.md)
