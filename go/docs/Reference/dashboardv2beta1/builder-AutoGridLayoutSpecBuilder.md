---
title: <span class="badge builder"></span> AutoGridLayoutSpecBuilder
---
# <span class="badge builder"></span> AutoGridLayoutSpecBuilder

## Constructor

```go
func NewAutoGridLayoutSpecBuilder() *AutoGridLayoutSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AutoGridLayoutSpecBuilder) Build() (AutoGridLayoutSpec, error)
```

### <span class="badge object-method"></span> ColumnWidth

```go
func (builder *AutoGridLayoutSpecBuilder) ColumnWidth(columnWidth float64) *AutoGridLayoutSpecBuilder
```

### <span class="badge object-method"></span> ColumnWidthMode

```go
func (builder *AutoGridLayoutSpecBuilder) ColumnWidthMode(columnWidthMode dashboardv2beta1.AutoGridLayoutSpecColumnWidthMode) *AutoGridLayoutSpecBuilder
```

### <span class="badge object-method"></span> FillScreen

```go
func (builder *AutoGridLayoutSpecBuilder) FillScreen(fillScreen bool) *AutoGridLayoutSpecBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *AutoGridLayoutSpecBuilder) Items(items []cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]) *AutoGridLayoutSpecBuilder
```

### <span class="badge object-method"></span> MaxColumnCount

```go
func (builder *AutoGridLayoutSpecBuilder) MaxColumnCount(maxColumnCount float64) *AutoGridLayoutSpecBuilder
```

### <span class="badge object-method"></span> RowHeight

```go
func (builder *AutoGridLayoutSpecBuilder) RowHeight(rowHeight float64) *AutoGridLayoutSpecBuilder
```

### <span class="badge object-method"></span> RowHeightMode

```go
func (builder *AutoGridLayoutSpecBuilder) RowHeightMode(rowHeightMode dashboardv2beta1.AutoGridLayoutSpecRowHeightMode) *AutoGridLayoutSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AutoGridLayoutSpec](./object-AutoGridLayoutSpec.md)
