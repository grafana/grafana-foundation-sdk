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
func (builder *PanelBuilder) Build() (dashboard.Panel, error)
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *PanelBuilder) AxisCenteredZero(axisCenteredZero bool) *PanelBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *PanelBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *PanelBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *PanelBuilder) AxisGridShow(axisGridShow bool) *PanelBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *PanelBuilder) AxisLabel(axisLabel string) *PanelBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *PanelBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *PanelBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *PanelBuilder) AxisSoftMax(axisSoftMax float64) *PanelBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *PanelBuilder) AxisSoftMin(axisSoftMin float64) *PanelBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *PanelBuilder) AxisWidth(axisWidth float64) *PanelBuilder
```

### <span class="badge object-method"></span> BarAlignment

```go
func (builder *PanelBuilder) BarAlignment(barAlignment common.BarAlignment) *PanelBuilder
```

### <span class="badge object-method"></span> BarMaxWidth

```go
func (builder *PanelBuilder) BarMaxWidth(barMaxWidth float64) *PanelBuilder
```

### <span class="badge object-method"></span> BarWidthFactor

```go
func (builder *PanelBuilder) BarWidthFactor(barWidthFactor float64) *PanelBuilder
```

### <span class="badge object-method"></span> CandleStyle

Sets the style of the candlesticks

```go
func (builder *PanelBuilder) CandleStyle(candleStyle candlestick.CandleStyle) *PanelBuilder
```

### <span class="badge object-method"></span> ColorScheme

Panel color configuration

```go
func (builder *PanelBuilder) ColorScheme(color cog.Builder[dashboard.FieldColor]) *PanelBuilder
```

### <span class="badge object-method"></span> ColorStrategy

Sets the color strategy for the candlesticks

```go
func (builder *PanelBuilder) ColorStrategy(colorStrategy candlestick.ColorStrategy) *PanelBuilder
```

### <span class="badge object-method"></span> Colors

Set which colors are used when the price movement is up or down

```go
func (builder *PanelBuilder) Colors(colors cog.Builder[candlestick.CandlestickColors]) *PanelBuilder
```

### <span class="badge object-method"></span> DataLinks

The behavior when clicking on a result

```go
func (builder *PanelBuilder) DataLinks(links []cog.Builder[dashboard.DashboardLink]) *PanelBuilder
```

### <span class="badge object-method"></span> Datasource

The datasource used in all targets.

```go
func (builder *PanelBuilder) Datasource(datasource dashboard.DataSourceRef) *PanelBuilder
```

### <span class="badge object-method"></span> Decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```go
func (builder *PanelBuilder) Decimals(decimals float64) *PanelBuilder
```

### <span class="badge object-method"></span> Description

Panel description.

```go
func (builder *PanelBuilder) Description(description string) *PanelBuilder
```

### <span class="badge object-method"></span> DisplayName

The display value for this field.  This supports template variables blank is auto

```go
func (builder *PanelBuilder) DisplayName(displayName string) *PanelBuilder
```

### <span class="badge object-method"></span> DrawStyle

```go
func (builder *PanelBuilder) DrawStyle(drawStyle common.GraphDrawStyle) *PanelBuilder
```

### <span class="badge object-method"></span> Fields

Map fields to appropriate dimension

```go
func (builder *PanelBuilder) Fields(fields cog.Builder[candlestick.CandlestickFieldMap]) *PanelBuilder
```

### <span class="badge object-method"></span> FillBelowTo

```go
func (builder *PanelBuilder) FillBelowTo(fillBelowTo string) *PanelBuilder
```

### <span class="badge object-method"></span> FillColor

```go
func (builder *PanelBuilder) FillColor(fillColor string) *PanelBuilder
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *PanelBuilder) FillOpacity(fillOpacity float64) *PanelBuilder
```

### <span class="badge object-method"></span> GradientMode

```go
func (builder *PanelBuilder) GradientMode(gradientMode common.GraphGradientMode) *PanelBuilder
```

### <span class="badge object-method"></span> GridPos

Grid position.

```go
func (builder *PanelBuilder) GridPos(gridPos dashboard.GridPos) *PanelBuilder
```

### <span class="badge object-method"></span> Height

Panel height. The height is the number of rows from the top edge of the panel.

```go
func (builder *PanelBuilder) Height(h uint32) *PanelBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *PanelBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *PanelBuilder
```

### <span class="badge object-method"></span> Id

Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.

```go
func (builder *PanelBuilder) Id(id uint32) *PanelBuilder
```

### <span class="badge object-method"></span> IncludeAllFields

When enabled, all fields will be sent to the graph

```go
func (builder *PanelBuilder) IncludeAllFields(includeAllFields bool) *PanelBuilder
```

### <span class="badge object-method"></span> InsertNulls

```go
func (builder *PanelBuilder) InsertNulls(insertNulls common.BoolOrUint32) *PanelBuilder
```

### <span class="badge object-method"></span> Interval

The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.

This value must be formatted as a number followed by a valid time

identifier like: "40s", "3d", etc.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *PanelBuilder) Interval(interval string) *PanelBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *PanelBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *PanelBuilder
```

### <span class="badge object-method"></span> LibraryPanel

Dynamically load the panel

```go
func (builder *PanelBuilder) LibraryPanel(libraryPanel dashboard.LibraryPanelRef) *PanelBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *PanelBuilder) LineColor(lineColor string) *PanelBuilder
```

### <span class="badge object-method"></span> LineInterpolation

```go
func (builder *PanelBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *PanelBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *PanelBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *PanelBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *PanelBuilder) LineWidth(lineWidth float64) *PanelBuilder
```

### <span class="badge object-method"></span> Links

Panel links.

```go
func (builder *PanelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *PanelBuilder
```

### <span class="badge object-method"></span> Mappings

Convert input values into a display string

```go
func (builder *PanelBuilder) Mappings(mappings []dashboard.ValueMapping) *PanelBuilder
```

### <span class="badge object-method"></span> Max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *PanelBuilder) Max(max float64) *PanelBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

The maximum number of data points that the panel queries are retrieving.

```go
func (builder *PanelBuilder) MaxDataPoints(maxDataPoints float64) *PanelBuilder
```

### <span class="badge object-method"></span> Min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *PanelBuilder) Min(min float64) *PanelBuilder
```

### <span class="badge object-method"></span> Mode

Sets which dimensions are used for the visualization

```go
func (builder *PanelBuilder) Mode(mode candlestick.VizDisplayMode) *PanelBuilder
```

### <span class="badge object-method"></span> NoValue

Alternative to empty string

```go
func (builder *PanelBuilder) NoValue(noValue string) *PanelBuilder
```

### <span class="badge object-method"></span> OverrideByFieldType

Adds override rules for all the fields of the given type.

```go
func (builder *PanelBuilder) OverrideByFieldType(fieldType string, properties []dashboard.DynamicConfigValue) *PanelBuilder
```

### <span class="badge object-method"></span> OverrideByName

Adds override rules for a specific field, referred to by its name.

```go
func (builder *PanelBuilder) OverrideByName(name string, properties []dashboard.DynamicConfigValue) *PanelBuilder
```

### <span class="badge object-method"></span> OverrideByQuery

```go
func (builder *PanelBuilder) OverrideByQuery(queryRefId string, properties []dashboard.DynamicConfigValue) *PanelBuilder
```

### <span class="badge object-method"></span> OverrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```go
func (builder *PanelBuilder) OverrideByRegexp(regexp string, properties []dashboard.DynamicConfigValue) *PanelBuilder
```

### <span class="badge object-method"></span> Overrides

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *PanelBuilder) Overrides(overrides []cog.Builder[dashboard.DashboardFieldConfigSourceOverrides]) *PanelBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *PanelBuilder) PointColor(pointColor string) *PanelBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *PanelBuilder) PointSize(pointSize float64) *PanelBuilder
```

### <span class="badge object-method"></span> PointSymbol

```go
func (builder *PanelBuilder) PointSymbol(pointSymbol string) *PanelBuilder
```

### <span class="badge object-method"></span> Repeat

Name of template variable to repeat for.

```go
func (builder *PanelBuilder) Repeat(repeat string) *PanelBuilder
```

### <span class="badge object-method"></span> RepeatDirection

Direction to repeat in if 'repeat' is set.

`h` for horizontal, `v` for vertical.

```go
func (builder *PanelBuilder) RepeatDirection(repeatDirection dashboard.PanelRepeatDirection) *PanelBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *PanelBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *PanelBuilder
```

### <span class="badge object-method"></span> ShowPoints

```go
func (builder *PanelBuilder) ShowPoints(showPoints common.VisibilityMode) *PanelBuilder
```

### <span class="badge object-method"></span> Span

Panel width. The width is the number of columns from the left edge of the panel.

```go
func (builder *PanelBuilder) Span(w uint32) *PanelBuilder
```

### <span class="badge object-method"></span> SpanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```go
func (builder *PanelBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *PanelBuilder
```

### <span class="badge object-method"></span> Stacking

```go
func (builder *PanelBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *PanelBuilder
```

### <span class="badge object-method"></span> Targets

Depends on the panel plugin. See the plugin documentation for details.

```go
func (builder *PanelBuilder) Targets(targets []cog.Builder[cog/variants.Dataquery]) *PanelBuilder
```

### <span class="badge object-method"></span> Thresholds

Map numeric values to states

```go
func (builder *PanelBuilder) Thresholds(thresholds cog.Builder[dashboard.ThresholdsConfig]) *PanelBuilder
```

### <span class="badge object-method"></span> ThresholdsStyle

```go
func (builder *PanelBuilder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *PanelBuilder
```

### <span class="badge object-method"></span> TimeFrom

Overrides the relative time range for individual panels,

which causes them to be different than what is selected in

the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different

time periods or days on the same dashboard.

The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),

`now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *PanelBuilder) TimeFrom(timeFrom string) *PanelBuilder
```

### <span class="badge object-method"></span> TimeShift

Overrides the time range for individual panels by shifting its start and end relative to the time picker.

For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *PanelBuilder) TimeShift(timeShift string) *PanelBuilder
```

### <span class="badge object-method"></span> Title

Panel title.

```go
func (builder *PanelBuilder) Title(title string) *PanelBuilder
```

### <span class="badge object-method"></span> Transform

```go
func (builder *PanelBuilder) Transform(transform common.GraphTransform) *PanelBuilder
```

### <span class="badge object-method"></span> Transformations

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```go
func (builder *PanelBuilder) Transformations(transformations []dashboard.DataTransformerConfig) *PanelBuilder
```

### <span class="badge object-method"></span> Transparent

Whether to display the panel without a background.

```go
func (builder *PanelBuilder) Transparent(transparent bool) *PanelBuilder
```

### <span class="badge object-method"></span> Unit

Unit a field should use. The unit you select is applied to all fields except time.

You can use the units ID availables in Grafana or a custom unit.

Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts

As custom unit, you can use the following formats:

`suffix:<suffix>` for custom unit that should go after value.

`prefix:<prefix>` for custom unit that should go before value.

`time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.

`si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.

`count:<unit>` for a custom count unit.

`currency:<unit>` for custom a currency unit.

```go
func (builder *PanelBuilder) Unit(unit string) *PanelBuilder
```

### <span class="badge object-method"></span> WithOverride

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *PanelBuilder) WithOverride(matcher dashboard.MatcherConfig, properties []dashboard.DynamicConfigValue) *PanelBuilder
```

### <span class="badge object-method"></span> WithTarget

Depends on the panel plugin. See the plugin documentation for details.

```go
func (builder *PanelBuilder) WithTarget(target cog.Builder[cog/variants.Dataquery]) *PanelBuilder
```

### <span class="badge object-method"></span> WithTransformation

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```go
func (builder *PanelBuilder) WithTransformation(transformation dashboard.DataTransformerConfig) *PanelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboard.Panel](../dashboard/object-Panel.md)
