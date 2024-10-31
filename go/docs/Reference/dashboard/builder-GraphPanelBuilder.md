---
title: <span class="badge builder"></span> GraphPanelBuilder
---
# <span class="badge builder"></span> GraphPanelBuilder

## Constructor

```go
func NewGraphPanelBuilder() *GraphPanelBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GraphPanelBuilder) Build() (GraphPanel, error)
```

### <span class="badge object-method"></span> Legend

@deprecated this is part of deprecated graph panel

```go
func (builder *GraphPanelBuilder) Legend(legend cog.Builder[dashboard.DashboardGraphPanelLegend]) *GraphPanelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GraphPanel](./object-GraphPanel.md)
