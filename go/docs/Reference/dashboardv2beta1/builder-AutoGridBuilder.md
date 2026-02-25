---
title: <span class="badge builder"></span> AutoGridBuilder
---
# <span class="badge builder"></span> AutoGridBuilder

## Constructor

```go
func NewAutoGridBuilder() *AutoGridBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AutoGridBuilder) Build() (AutoGridLayoutKind, error)
```

### <span class="badge object-method"></span> ColumnWidth

```go
func (builder *AutoGridBuilder) ColumnWidth(columnWidth float64) *AutoGridBuilder
```

### <span class="badge object-method"></span> ColumnWidthMode

```go
func (builder *AutoGridBuilder) ColumnWidthMode(columnWidthMode dashboardv2beta1.AutoGridLayoutSpecColumnWidthMode) *AutoGridBuilder
```

### <span class="badge object-method"></span> FillScreen

```go
func (builder *AutoGridBuilder) FillScreen(fillScreen bool) *AutoGridBuilder
```

### <span class="badge object-method"></span> Item

```go
func (builder *AutoGridBuilder) Item(item cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]) *AutoGridBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *AutoGridBuilder) Items(items []cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]) *AutoGridBuilder
```

### <span class="badge object-method"></span> MaxColumnCount

```go
func (builder *AutoGridBuilder) MaxColumnCount(maxColumnCount float64) *AutoGridBuilder
```

### <span class="badge object-method"></span> RowHeight

```go
func (builder *AutoGridBuilder) RowHeight(rowHeight float64) *AutoGridBuilder
```

### <span class="badge object-method"></span> RowHeightMode

```go
func (builder *AutoGridBuilder) RowHeightMode(rowHeightMode dashboardv2beta1.AutoGridLayoutSpecRowHeightMode) *AutoGridBuilder
```

### <span class="badge object-method"></span> WithItem

```go
func (builder *AutoGridBuilder) WithItem(name string) *AutoGridBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AutoGridLayoutKind](./object-AutoGridLayoutKind.md)
