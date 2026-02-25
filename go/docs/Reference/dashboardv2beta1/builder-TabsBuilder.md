---
title: <span class="badge builder"></span> TabsBuilder
---
# <span class="badge builder"></span> TabsBuilder

## Constructor

```go
func NewTabsBuilder() *TabsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TabsBuilder) Build() (TabsLayoutKind, error)
```

### <span class="badge object-method"></span> Tab

```go
func (builder *TabsBuilder) Tab(tab cog.Builder[dashboardv2beta1.TabsLayoutTabKind]) *TabsBuilder
```

### <span class="badge object-method"></span> Tabs

```go
func (builder *TabsBuilder) Tabs(tabs []cog.Builder[dashboardv2beta1.TabsLayoutTabKind]) *TabsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TabsLayoutKind](./object-TabsLayoutKind.md)
