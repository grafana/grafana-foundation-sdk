---
title: <span class="badge builder"></span> AutoGridLayoutBuilder
---
# <span class="badge builder"></span> AutoGridLayoutBuilder

## Constructor

```go
func NewAutoGridLayoutBuilder() *AutoGridLayoutBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AutoGridLayoutBuilder) Build() (AutoGridLayoutKind, error)
```

### <span class="badge object-method"></span> ColumnWidth

```go
func (builder *AutoGridLayoutBuilder) ColumnWidth(columnWidth float64) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> ColumnWidthMode

```go
func (builder *AutoGridLayoutBuilder) ColumnWidthMode(columnWidthMode dashboardv2beta1.AutoGridLayoutSpecColumnWidthMode) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> FillScreen

```go
func (builder *AutoGridLayoutBuilder) FillScreen(fillScreen bool) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> Item

```go
func (builder *AutoGridLayoutBuilder) Item(item cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *AutoGridLayoutBuilder) Items(items []cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> MaxColumnCount

```go
func (builder *AutoGridLayoutBuilder) MaxColumnCount(maxColumnCount float64) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> RowHeight

```go
func (builder *AutoGridLayoutBuilder) RowHeight(rowHeight float64) *AutoGridLayoutBuilder
```

### <span class="badge object-method"></span> RowHeightMode

```go
func (builder *AutoGridLayoutBuilder) RowHeightMode(rowHeightMode dashboardv2beta1.AutoGridLayoutSpecRowHeightMode) *AutoGridLayoutBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AutoGridLayoutKind](./object-AutoGridLayoutKind.md)
