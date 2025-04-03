---
title: <span class="badge builder"></span> NestedBuilder
---
# <span class="badge builder"></span> NestedBuilder

## Constructor

```go
func NewNestedBuilder() *NestedBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NestedBuilder) Build() (Nested, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *NestedBuilder) Field(field string) *NestedBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *NestedBuilder) Id(id string) *NestedBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *NestedBuilder) Settings(settings any) *NestedBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Nested](./object-Nested.md)
