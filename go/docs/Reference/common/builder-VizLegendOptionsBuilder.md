---
title: <span class="badge builder"></span> VizLegendOptionsBuilder
---
# <span class="badge builder"></span> VizLegendOptionsBuilder

## Constructor

```go
func NewVizLegendOptionsBuilder() *VizLegendOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VizLegendOptionsBuilder) Build() (VizLegendOptions, error)
```

### <span class="badge object-method"></span> AsTable

```go
func (builder *VizLegendOptionsBuilder) AsTable(asTable bool) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> Calcs

```go
func (builder *VizLegendOptionsBuilder) Calcs(calcs []string) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> DisplayMode

```go
func (builder *VizLegendOptionsBuilder) DisplayMode(displayMode common.LegendDisplayMode) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> IsVisible

```go
func (builder *VizLegendOptionsBuilder) IsVisible(isVisible bool) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> Placement

```go
func (builder *VizLegendOptionsBuilder) Placement(placement common.LegendPlacement) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> ShowLegend

```go
func (builder *VizLegendOptionsBuilder) ShowLegend(showLegend bool) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> SortBy

```go
func (builder *VizLegendOptionsBuilder) SortBy(sortBy string) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> SortDesc

```go
func (builder *VizLegendOptionsBuilder) SortDesc(sortDesc bool) *VizLegendOptionsBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *VizLegendOptionsBuilder) Width(width float64) *VizLegendOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VizLegendOptions](./object-VizLegendOptions.md)
