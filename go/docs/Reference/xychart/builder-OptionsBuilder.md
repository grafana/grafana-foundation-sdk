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

### <span class="badge object-method"></span> Dims

Table Mode (auto)

```go
func (builder *OptionsBuilder) Dims(dims cog.Builder[xychart.XYDimensionConfig]) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Series

Manual Mode

```go
func (builder *OptionsBuilder) Series(series []cog.Builder[xychart.ScatterSeriesConfig]) *OptionsBuilder
```

### <span class="badge object-method"></span> SeriesMapping

```go
func (builder *OptionsBuilder) SeriesMapping(seriesMapping xychart.SeriesMapping) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
