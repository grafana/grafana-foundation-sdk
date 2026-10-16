---
title: <span class="badge builder"></span> V2PanelSpecBuilder
---
# <span class="badge builder"></span> V2PanelSpecBuilder

## Constructor

```go
func NewV2PanelSpecBuilder() *V2PanelSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *V2PanelSpecBuilder) Build() (V2PanelSpec, error)
```

### <span class="badge object-method"></span> Data

```go
func (builder *V2PanelSpecBuilder) Data(data cog.Builder[dashboardv2beta1.V2QueryGroupKind]) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> Description

Shown in a info icon tooltip next to panel title

```go
func (builder *V2PanelSpecBuilder) Description(description string) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *V2PanelSpecBuilder) Id(id float64) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> Links

```go
func (builder *V2PanelSpecBuilder) Links(links []cog.Builder[dashboardv2beta1.DataLink]) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> Subtitle

Shown in a sub header below the title.

```go
func (builder *V2PanelSpecBuilder) Subtitle(subtitle string) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *V2PanelSpecBuilder) Title(title string) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> Transparent

```go
func (builder *V2PanelSpecBuilder) Transparent(transparent bool) *V2PanelSpecBuilder
```

### <span class="badge object-method"></span> VizConfig

```go
func (builder *V2PanelSpecBuilder) VizConfig(vizConfig cog.Builder[dashboardv2beta1.VizConfigKind]) *V2PanelSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [V2PanelSpec](./object-V2PanelSpec.md)
