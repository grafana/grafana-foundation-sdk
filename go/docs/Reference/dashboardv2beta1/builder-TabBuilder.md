---
title: <span class="badge builder"></span> TabBuilder
---
# <span class="badge builder"></span> TabBuilder

## Constructor

```go
func NewTabBuilder() *TabBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TabBuilder) Build() (TabsLayoutTabKind, error)
```

### <span class="badge object-method"></span> AutoGridLayout

```go
func (builder *TabBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[dashboardv2beta1.AutoGridLayoutKind]) *TabBuilder
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *TabBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *TabBuilder
```

### <span class="badge object-method"></span> GridLayout

```go
func (builder *TabBuilder) GridLayout(gridLayoutKind cog.Builder[dashboardv2beta1.GridLayoutKind]) *TabBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *TabBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.TabRepeatOptions]) *TabBuilder
```

### <span class="badge object-method"></span> RowsLayout

```go
func (builder *TabBuilder) RowsLayout(rowsLayoutKind cog.Builder[dashboardv2beta1.RowsLayoutKind]) *TabBuilder
```

### <span class="badge object-method"></span> TabsLayout

```go
func (builder *TabBuilder) TabsLayout(tabsLayoutKind cog.Builder[dashboardv2beta1.TabsLayoutKind]) *TabBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *TabBuilder) Title(title string) *TabBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TabsLayoutTabKind](./object-TabsLayoutTabKind.md)
