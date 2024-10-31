---
title: <span class="badge object-type-struct"></span> QueryEditorOperator
---
# <span class="badge object-type-struct"></span> QueryEditorOperator

TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer

## Definition

```go
type QueryEditorOperator struct {
    Name *string `json:"name,omitempty"`
    Value *cloudwatch.StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType `json:"value,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorOperator` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorOperator *QueryEditorOperator) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorOperator` objects.

```go
func (queryEditorOperator *QueryEditorOperator) Equals(other QueryEditorOperator) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorOperator` fields for violations and returns them.

```go
func (queryEditorOperator *QueryEditorOperator) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorOperatorBuilder](./builder-QueryEditorOperatorBuilder.md)
