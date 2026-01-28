---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> DisplayLabels

```go
func (builder *OptionsBuilder) DisplayLabels(displayLabels []piechart.PieChartLabels) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[piechart.PieChartLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Orientation

```go
func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder
```

### <span class="badge object-method"></span> PieType

```go
func (builder *OptionsBuilder) PieType(pieType piechart.PieChartType) *OptionsBuilder
```

### <span class="badge object-method"></span> ReduceOptions

```go
func (builder *OptionsBuilder) ReduceOptions(reduceOptions cog.Builder[common.ReduceDataOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
