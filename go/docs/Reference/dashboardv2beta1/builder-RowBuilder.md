---
title: <span class="badge builder"></span> RowBuilder
---
# <span class="badge builder"></span> RowBuilder

## Constructor

```go
func NewRowBuilder() *RowBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowBuilder) Build() (RowsLayoutRowKind, error)
```

### <span class="badge object-method"></span> AutoGridLayout

```go
func (builder *RowBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[dashboardv2beta1.AutoGridLayoutKind]) *RowBuilder
```

### <span class="badge object-method"></span> Collapse

```go
func (builder *RowBuilder) Collapse(collapse bool) *RowBuilder
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *RowBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *RowBuilder
```

### <span class="badge object-method"></span> FillScreen

```go
func (builder *RowBuilder) FillScreen(fillScreen bool) *RowBuilder
```

### <span class="badge object-method"></span> GridLayout

```go
func (builder *RowBuilder) GridLayout(gridLayoutKind cog.Builder[dashboardv2beta1.GridLayoutKind]) *RowBuilder
```

### <span class="badge object-method"></span> HideHeader

```go
func (builder *RowBuilder) HideHeader(hideHeader bool) *RowBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *RowBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.RowRepeatOptions]) *RowBuilder
```

### <span class="badge object-method"></span> RowsLayout

```go
func (builder *RowBuilder) RowsLayout(rowsLayoutKind cog.Builder[dashboardv2beta1.RowsLayoutKind]) *RowBuilder
```

### <span class="badge object-method"></span> TabsLayout

```go
func (builder *RowBuilder) TabsLayout(tabsLayoutKind cog.Builder[dashboardv2beta1.TabsLayoutKind]) *RowBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RowBuilder) Title(title string) *RowBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsLayoutRowKind](./object-RowsLayoutRowKind.md)
