---
title: <span class="badge builder"></span> TabsLayoutBuilder
---
# <span class="badge builder"></span> TabsLayoutBuilder

## Constructor

```go
func NewTabsLayoutBuilder() *TabsLayoutBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TabsLayoutBuilder) Build() (TabsLayoutKind, error)
```

### <span class="badge object-method"></span> Tab

```go
func (builder *TabsLayoutBuilder) Tab(tab cog.Builder[dashboardv2beta1.TabsLayoutTabKind]) *TabsLayoutBuilder
```

### <span class="badge object-method"></span> Tabs

```go
func (builder *TabsLayoutBuilder) Tabs(tabs []cog.Builder[dashboardv2beta1.TabsLayoutTabKind]) *TabsLayoutBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TabsLayoutKind](./object-TabsLayoutKind.md)
