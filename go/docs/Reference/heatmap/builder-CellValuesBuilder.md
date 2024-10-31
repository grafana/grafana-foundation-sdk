---
title: <span class="badge builder"></span> CellValuesBuilder
---
# <span class="badge builder"></span> CellValuesBuilder

## Constructor

```go
func NewCellValuesBuilder() *CellValuesBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CellValuesBuilder) Build() (CellValues, error)
```

### <span class="badge object-method"></span> Decimals

Controls the number of decimals for cell values

```go
func (builder *CellValuesBuilder) Decimals(decimals float32) *CellValuesBuilder
```

### <span class="badge object-method"></span> Unit

Controls the cell value unit

```go
func (builder *CellValuesBuilder) Unit(unit string) *CellValuesBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CellValues](./object-CellValues.md)
