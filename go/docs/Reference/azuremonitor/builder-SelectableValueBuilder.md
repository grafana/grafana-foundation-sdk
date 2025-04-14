---
title: <span class="badge builder"></span> SelectableValueBuilder
---
# <span class="badge builder"></span> SelectableValueBuilder

## Constructor

```go
func NewSelectableValueBuilder() *SelectableValueBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SelectableValueBuilder) Build() (SelectableValue, error)
```

### <span class="badge object-method"></span> Label

```go
func (builder *SelectableValueBuilder) Label(label string) *SelectableValueBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *SelectableValueBuilder) Value(value string) *SelectableValueBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SelectableValue](./object-SelectableValue.md)
