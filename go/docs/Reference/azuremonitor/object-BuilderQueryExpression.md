---
title: <span class="badge object-type-struct"></span> BuilderQueryExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryExpression

## Definition

```go
type BuilderQueryExpression struct {
    From *azuremonitor.BuilderQueryEditorPropertyExpression `json:"from,omitempty"`
    Columns *azuremonitor.BuilderQueryEditorColumnsExpression `json:"columns,omitempty"`
    Where *azuremonitor.BuilderQueryEditorWhereExpressionArray `json:"where,omitempty"`
    Reduce *azuremonitor.BuilderQueryEditorReduceExpressionArray `json:"reduce,omitempty"`
    GroupBy *azuremonitor.BuilderQueryEditorGroupByExpressionArray `json:"groupBy,omitempty"`
    Limit *int64 `json:"limit,omitempty"`
    OrderBy *azuremonitor.BuilderQueryEditorOrderByExpressionArray `json:"orderBy,omitempty"`
    FuzzySearch *azuremonitor.BuilderQueryEditorWhereExpressionArray `json:"fuzzySearch,omitempty"`
    TimeFilter *azuremonitor.BuilderQueryEditorWhereExpressionArray `json:"timeFilter,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryExpression *BuilderQueryExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryExpression` objects.

```go
func (builderQueryExpression *BuilderQueryExpression) Equals(other BuilderQueryExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryExpression` fields for violations and returns them.

```go
func (builderQueryExpression *BuilderQueryExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryExpressionBuilder](./builder-BuilderQueryExpressionBuilder.md)
