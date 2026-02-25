---
title: <span class="badge builder"></span> RowsBuilder
---
# <span class="badge builder"></span> RowsBuilder

## Constructor

```go
func NewRowsBuilder() *RowsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowsBuilder) Build() (RowsLayoutKind, error)
```

### <span class="badge object-method"></span> Row

```go
func (builder *RowsBuilder) Row(row cog.Builder[dashboardv2beta1.RowsLayoutRowKind]) *RowsBuilder
```

### <span class="badge object-method"></span> Rows

```go
func (builder *RowsBuilder) Rows(rows []cog.Builder[dashboardv2beta1.RowsLayoutRowKind]) *RowsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsLayoutKind](./object-RowsLayoutKind.md)
