---
title: <span class="badge builder"></span> TabsLayoutTabSpecBuilder
---
# <span class="badge builder"></span> TabsLayoutTabSpecBuilder

## Constructor

```go
func NewTabsLayoutTabSpecBuilder() *TabsLayoutTabSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TabsLayoutTabSpecBuilder) Build() (TabsLayoutTabSpec, error)
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *TabsLayoutTabSpecBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *TabsLayoutTabSpecBuilder
```

### <span class="badge object-method"></span> Layout

```go
func (builder *TabsLayoutTabSpecBuilder) Layout(layout dashboardv2beta1.GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) *TabsLayoutTabSpecBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *TabsLayoutTabSpecBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.TabRepeatOptions]) *TabsLayoutTabSpecBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *TabsLayoutTabSpecBuilder) Title(title string) *TabsLayoutTabSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TabsLayoutTabSpec](./object-TabsLayoutTabSpec.md)
