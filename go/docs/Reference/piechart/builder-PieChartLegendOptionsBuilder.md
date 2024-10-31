---
title: <span class="badge builder"></span> PieChartLegendOptionsBuilder
---
# <span class="badge builder"></span> PieChartLegendOptionsBuilder

## Constructor

```go
func NewPieChartLegendOptionsBuilder() *PieChartLegendOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PieChartLegendOptionsBuilder) Build() (PieChartLegendOptions, error)
```

### <span class="badge object-method"></span> AsTable

```go
func (builder *PieChartLegendOptionsBuilder) AsTable(asTable bool) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> Calcs

```go
func (builder *PieChartLegendOptionsBuilder) Calcs(calcs []string) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> DisplayMode

```go
func (builder *PieChartLegendOptionsBuilder) DisplayMode(displayMode common.LegendDisplayMode) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> IsVisible

```go
func (builder *PieChartLegendOptionsBuilder) IsVisible(isVisible bool) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> Placement

```go
func (builder *PieChartLegendOptionsBuilder) Placement(placement common.LegendPlacement) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> ShowLegend

```go
func (builder *PieChartLegendOptionsBuilder) ShowLegend(showLegend bool) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> SortBy

```go
func (builder *PieChartLegendOptionsBuilder) SortBy(sortBy string) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> SortDesc

```go
func (builder *PieChartLegendOptionsBuilder) SortDesc(sortDesc bool) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> Values

```go
func (builder *PieChartLegendOptionsBuilder) Values(values []piechart.PieChartLegendValues) *PieChartLegendOptionsBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *PieChartLegendOptionsBuilder) Width(width float64) *PieChartLegendOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PieChartLegendOptions](./object-PieChartLegendOptions.md)
