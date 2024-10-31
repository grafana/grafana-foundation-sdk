---
title: <span class="badge object-type-struct"></span> QueryEditorOperatorExpression
---
# <span class="badge object-type-struct"></span> QueryEditorOperatorExpression

## Definition

```go
type QueryEditorOperatorExpression struct {
    Type string `json:"type"`
    Property cloudwatch.QueryEditorProperty `json:"property"`
    // TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    Operator cloudwatch.QueryEditorOperator `json:"operator"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorOperatorExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorOperatorExpression *QueryEditorOperatorExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorOperatorExpression` objects.

```go
func (queryEditorOperatorExpression *QueryEditorOperatorExpression) Equals(other QueryEditorOperatorExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorOperatorExpression` fields for violations and returns them.

```go
func (queryEditorOperatorExpression *QueryEditorOperatorExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorOperatorExpressionBuilder](./builder-QueryEditorOperatorExpressionBuilder.md)
