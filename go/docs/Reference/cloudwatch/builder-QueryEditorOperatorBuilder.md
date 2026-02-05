---
title: <span class="badge builder"></span> QueryEditorOperatorBuilder
---
# <span class="badge builder"></span> QueryEditorOperatorBuilder

## Constructor

```go
func NewQueryEditorOperatorBuilder() *QueryEditorOperatorBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryEditorOperatorBuilder) Build() (QueryEditorOperator, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *QueryEditorOperatorBuilder) Name(name string) *QueryEditorOperatorBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *QueryEditorOperatorBuilder) Value(operatorTypes []cloudwatch.QueryEditorOperatorType) *QueryEditorOperatorBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryEditorOperator](./object-QueryEditorOperator.md)
