---
title: <span class="badge builder"></span> RowsHeatmapOptionsBuilder
---
# <span class="badge builder"></span> RowsHeatmapOptionsBuilder

## Constructor

```go
func NewRowsHeatmapOptionsBuilder() *RowsHeatmapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowsHeatmapOptionsBuilder) Build() (RowsHeatmapOptions, error)
```

### <span class="badge object-method"></span> Layout

Controls tick alignment when not calculating from data

```go
func (builder *RowsHeatmapOptionsBuilder) Layout(layout common.HeatmapCellLayout) *RowsHeatmapOptionsBuilder
```

### <span class="badge object-method"></span> Value

Sets the name of the cell when not calculating from data

```go
func (builder *RowsHeatmapOptionsBuilder) Value(value string) *RowsHeatmapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsHeatmapOptions](./object-RowsHeatmapOptions.md)
