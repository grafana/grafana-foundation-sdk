---
title: <span class="badge builder"></span> TabsLayoutTabBuilder
---
# <span class="badge builder"></span> TabsLayoutTabBuilder

## Constructor

```go
func NewTabsLayoutTabBuilder(title string) *TabsLayoutTabBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TabsLayoutTabBuilder) Build() (TabsLayoutTabKind, error)
```

### <span class="badge object-method"></span> AutoGridLayout

```go
func (builder *TabsLayoutTabBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[dashboardv2beta1.AutoGridLayoutKind]) *TabsLayoutTabBuilder
```

### <span class="badge object-method"></span> RowsLayout

```go
func (builder *TabsLayoutTabBuilder) RowsLayout(rowsLayoutKind cog.Builder[dashboardv2beta1.RowsLayoutKind]) *TabsLayoutTabBuilder
```

### <span class="badge object-method"></span> TabsLayout

```go
func (builder *TabsLayoutTabBuilder) TabsLayout(tabsLayoutKind cog.Builder[dashboardv2beta1.TabsLayoutKind]) *TabsLayoutTabBuilder
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *TabsLayoutTabBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *TabsLayoutTabBuilder
```

### <span class="badge object-method"></span> GridLayout

```go
func (builder *TabsLayoutTabBuilder) GridLayout(gridLayoutKind cog.Builder[dashboardv2beta1.GridLayoutKind]) *TabsLayoutTabBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *TabsLayoutTabBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.TabRepeatOptions]) *TabsLayoutTabBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *TabsLayoutTabBuilder) Title(title string) *TabsLayoutTabBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TabsLayoutTabKind](./object-TabsLayoutTabKind.md)
