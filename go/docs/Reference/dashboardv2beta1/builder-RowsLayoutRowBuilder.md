---
title: <span class="badge builder"></span> RowsLayoutRowBuilder
---
# <span class="badge builder"></span> RowsLayoutRowBuilder

## Constructor

```go
func NewRowsLayoutRowBuilder(title string) *RowsLayoutRowBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowsLayoutRowBuilder) Build() (RowsLayoutRowKind, error)
```

### <span class="badge object-method"></span> AutoGridLayout

```go
func (builder *RowsLayoutRowBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[dashboardv2beta1.AutoGridLayoutKind]) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> RowsLayout

```go
func (builder *RowsLayoutRowBuilder) RowsLayout(rowsLayoutKind cog.Builder[dashboardv2beta1.RowsLayoutKind]) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> TabsLayout

```go
func (builder *RowsLayoutRowBuilder) TabsLayout(tabsLayoutKind cog.Builder[dashboardv2beta1.TabsLayoutKind]) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> Collapse

```go
func (builder *RowsLayoutRowBuilder) Collapse(collapse bool) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *RowsLayoutRowBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> FillScreen

```go
func (builder *RowsLayoutRowBuilder) FillScreen(fillScreen bool) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> GridLayout

```go
func (builder *RowsLayoutRowBuilder) GridLayout(gridLayoutKind cog.Builder[dashboardv2beta1.GridLayoutKind]) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> HideHeader

```go
func (builder *RowsLayoutRowBuilder) HideHeader(hideHeader bool) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *RowsLayoutRowBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.RowRepeatOptions]) *RowsLayoutRowBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RowsLayoutRowBuilder) Title(title string) *RowsLayoutRowBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsLayoutRowKind](./object-RowsLayoutRowKind.md)
