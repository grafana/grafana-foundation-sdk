---
title: <span class="badge builder"></span> QueryEditorPropertyBuilder
---
# <span class="badge builder"></span> QueryEditorPropertyBuilder

## Constructor

```go
func NewQueryEditorPropertyBuilder() *QueryEditorPropertyBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryEditorPropertyBuilder) Build() (QueryEditorProperty, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *QueryEditorPropertyBuilder) Name(name string) *QueryEditorPropertyBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *QueryEditorPropertyBuilder) Type(typeArg cloudwatch.QueryEditorPropertyType) *QueryEditorPropertyBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryEditorProperty](./object-QueryEditorProperty.md)
