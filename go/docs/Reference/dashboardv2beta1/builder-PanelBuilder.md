---
title: <span class="badge builder"></span> PanelBuilder
---
# <span class="badge builder"></span> PanelBuilder

## Constructor

```go
func NewPanelBuilder() *PanelBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PanelBuilder) Build() (PanelKind, error)
```

### <span class="badge object-method"></span> Data

```go
func (builder *PanelBuilder) Data(data cog.Builder[dashboardv2beta1.QueryGroupKind]) *PanelBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *PanelBuilder) Description(description string) *PanelBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *PanelBuilder) Id(id float64) *PanelBuilder
```

### <span class="badge object-method"></span> Links

```go
func (builder *PanelBuilder) Links(links []cog.Builder[dashboardv2beta1.DataLink]) *PanelBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *PanelBuilder) Title(title string) *PanelBuilder
```

### <span class="badge object-method"></span> Transparent

```go
func (builder *PanelBuilder) Transparent(transparent bool) *PanelBuilder
```

### <span class="badge object-method"></span> Visualization

```go
func (builder *PanelBuilder) Visualization(vizConfig cog.Builder[dashboardv2beta1.VizConfigKind]) *PanelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PanelKind](./object-PanelKind.md)
