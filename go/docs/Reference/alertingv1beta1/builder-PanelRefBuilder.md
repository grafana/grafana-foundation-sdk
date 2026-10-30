---
title: <span class="badge builder"></span> PanelRefBuilder
---
# <span class="badge builder"></span> PanelRefBuilder

## Constructor

```go
func NewPanelRefBuilder() *PanelRefBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PanelRefBuilder) Build() (PanelRef, error)
```

### <span class="badge object-method"></span> DashboardUID

```go
func (builder *PanelRefBuilder) DashboardUID(dashboardUID string) *PanelRefBuilder
```

### <span class="badge object-method"></span> PanelID

```go
func (builder *PanelRefBuilder) PanelID(panelID int64) *PanelRefBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PanelRef](./object-PanelRef.md)
