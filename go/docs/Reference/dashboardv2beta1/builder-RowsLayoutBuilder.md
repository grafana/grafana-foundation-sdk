---
title: <span class="badge builder"></span> RowsLayoutBuilder
---
# <span class="badge builder"></span> RowsLayoutBuilder

## Constructor

```go
func NewRowsLayoutBuilder() *RowsLayoutBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowsLayoutBuilder) Build() (RowsLayoutKind, error)
```

### <span class="badge object-method"></span> Row

```go
func (builder *RowsLayoutBuilder) Row(row cog.Builder[dashboardv2beta1.RowsLayoutRowKind]) *RowsLayoutBuilder
```

### <span class="badge object-method"></span> Rows

```go
func (builder *RowsLayoutBuilder) Rows(rows []cog.Builder[dashboardv2beta1.RowsLayoutRowKind]) *RowsLayoutBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsLayoutKind](./object-RowsLayoutKind.md)
