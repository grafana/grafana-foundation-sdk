---
title: <span class="badge builder"></span> RowsLayoutSpecBuilder
---
# <span class="badge builder"></span> RowsLayoutSpecBuilder

## Constructor

```go
func NewRowsLayoutSpecBuilder() *RowsLayoutSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowsLayoutSpecBuilder) Build() (RowsLayoutSpec, error)
```

### <span class="badge object-method"></span> Rows

```go
func (builder *RowsLayoutSpecBuilder) Rows(rows []cog.Builder[dashboardv2beta1.RowsLayoutRowKind]) *RowsLayoutSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsLayoutSpec](./object-RowsLayoutSpec.md)
