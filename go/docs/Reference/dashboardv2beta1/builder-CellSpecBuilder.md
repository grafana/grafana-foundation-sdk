---
title: <span class="badge builder"></span> CellSpecBuilder
---
# <span class="badge builder"></span> CellSpecBuilder

## Constructor

```go
func NewCellSpecBuilder() *CellSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CellSpecBuilder) Build() (CellSpec, error)
```

### <span class="badge object-method"></span> Content

```go
func (builder *CellSpecBuilder) Content(content dashboardv2beta1.CellContentKind) *CellSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CellSpec](./object-CellSpec.md)
