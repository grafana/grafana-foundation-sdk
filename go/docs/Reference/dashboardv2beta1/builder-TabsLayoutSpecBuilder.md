---
title: <span class="badge builder"></span> TabsLayoutSpecBuilder
---
# <span class="badge builder"></span> TabsLayoutSpecBuilder

## Constructor

```go
func NewTabsLayoutSpecBuilder() *TabsLayoutSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TabsLayoutSpecBuilder) Build() (TabsLayoutSpec, error)
```

### <span class="badge object-method"></span> Tabs

```go
func (builder *TabsLayoutSpecBuilder) Tabs(tabs []cog.Builder[dashboardv2beta1.TabsLayoutTabKind]) *TabsLayoutSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TabsLayoutSpec](./object-TabsLayoutSpec.md)
