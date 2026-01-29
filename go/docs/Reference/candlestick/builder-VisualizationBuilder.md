---
title: <span class="badge builder"></span> VisualizationBuilder
---
# <span class="badge builder"></span> VisualizationBuilder

## Constructor

```go
func NewVisualizationBuilder() *VisualizationBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VisualizationBuilder) Build() (dashboardv2beta1.VizConfigKind, error)
```

### <span class="badge object-method"></span> Actions

Define interactive HTTP requests that can be triggered from data visualizations.

```go
func (builder *VisualizationBuilder) Actions(actions []cog.Builder[dashboardv2beta1.Action]) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisBorderShow

```go
func (builder *VisualizationBuilder) AxisBorderShow(axisBorderShow bool) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *VisualizationBuilder) AxisCenteredZero(axisCenteredZero bool) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *VisualizationBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *VisualizationBuilder) AxisGridShow(axisGridShow bool) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *VisualizationBuilder) AxisLabel(axisLabel string) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *VisualizationBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *VisualizationBuilder) AxisSoftMax(axisSoftMax float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *VisualizationBuilder) AxisSoftMin(axisSoftMin float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *VisualizationBuilder) AxisWidth(axisWidth float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> BarAlignment

```go
func (builder *VisualizationBuilder) BarAlignment(barAlignment common.BarAlignment) *VisualizationBuilder
```

### <span class="badge object-method"></span> BarMaxWidth

```go
func (builder *VisualizationBuilder) BarMaxWidth(barMaxWidth float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> BarWidthFactor

```go
func (builder *VisualizationBuilder) BarWidthFactor(barWidthFactor float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> CandleStyle

Sets the style of the candlesticks

```go
func (builder *VisualizationBuilder) CandleStyle(candleStyle candlestick.CandleStyle) *VisualizationBuilder
```

### <span class="badge object-method"></span> ColorScheme

Panel color configuration

```go
func (builder *VisualizationBuilder) ColorScheme(color cog.Builder[dashboardv2beta1.FieldColor]) *VisualizationBuilder
```

### <span class="badge object-method"></span> ColorStrategy

Sets the color strategy for the candlesticks

```go
func (builder *VisualizationBuilder) ColorStrategy(colorStrategy candlestick.ColorStrategy) *VisualizationBuilder
```

### <span class="badge object-method"></span> Colors

Set which colors are used when the price movement is up or down

```go
func (builder *VisualizationBuilder) Colors(colors cog.Builder[candlestick.CandlestickColors]) *VisualizationBuilder
```

### <span class="badge object-method"></span> DataLinks

The behavior when clicking on a result

```go
func (builder *VisualizationBuilder) DataLinks(links []any) *VisualizationBuilder
```

### <span class="badge object-method"></span> Decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```go
func (builder *VisualizationBuilder) Decimals(decimals float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Description

Human readable field metadata

```go
func (builder *VisualizationBuilder) Description(description string) *VisualizationBuilder
```

### <span class="badge object-method"></span> DisplayName

The display value for this field.  This supports template variables blank is auto

```go
func (builder *VisualizationBuilder) DisplayName(displayName string) *VisualizationBuilder
```

### <span class="badge object-method"></span> DisplayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```go
func (builder *VisualizationBuilder) DisplayNameFromDS(displayNameFromDS string) *VisualizationBuilder
```

### <span class="badge object-method"></span> DrawStyle

```go
func (builder *VisualizationBuilder) DrawStyle(drawStyle common.GraphDrawStyle) *VisualizationBuilder
```

### <span class="badge object-method"></span> FieldMinMax

Calculate min max per field

```go
func (builder *VisualizationBuilder) FieldMinMax(fieldMinMax bool) *VisualizationBuilder
```

### <span class="badge object-method"></span> Fields

Map fields to appropriate dimension

```go
func (builder *VisualizationBuilder) Fields(fields cog.Builder[candlestick.CandlestickFieldMap]) *VisualizationBuilder
```

### <span class="badge object-method"></span> FillBelowTo

```go
func (builder *VisualizationBuilder) FillBelowTo(fillBelowTo string) *VisualizationBuilder
```

### <span class="badge object-method"></span> FillColor

```go
func (builder *VisualizationBuilder) FillColor(fillColor string) *VisualizationBuilder
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *VisualizationBuilder) FillOpacity(fillOpacity float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> GradientMode

```go
func (builder *VisualizationBuilder) GradientMode(gradientMode common.GraphGradientMode) *VisualizationBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *VisualizationBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *VisualizationBuilder
```

### <span class="badge object-method"></span> IncludeAllFields

When enabled, all fields will be sent to the graph

```go
func (builder *VisualizationBuilder) IncludeAllFields(includeAllFields bool) *VisualizationBuilder
```

### <span class="badge object-method"></span> InsertNulls

```go
func (builder *VisualizationBuilder) InsertNulls(insertNulls common.BoolOrFloat64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *VisualizationBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *VisualizationBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *VisualizationBuilder) LineColor(lineColor string) *VisualizationBuilder
```

### <span class="badge object-method"></span> LineInterpolation

```go
func (builder *VisualizationBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *VisualizationBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *VisualizationBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *VisualizationBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *VisualizationBuilder) LineWidth(lineWidth float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Mappings

Convert input values into a display string

```go
func (builder *VisualizationBuilder) Mappings(mappings []dashboardv2beta1.ValueMapping) *VisualizationBuilder
```

### <span class="badge object-method"></span> Max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VisualizationBuilder) Max(max float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VisualizationBuilder) Min(min float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Mode

Sets which dimensions are used for the visualization

```go
func (builder *VisualizationBuilder) Mode(mode candlestick.VizDisplayMode) *VisualizationBuilder
```

### <span class="badge object-method"></span> NoValue

Alternative to empty string

```go
func (builder *VisualizationBuilder) NoValue(noValue string) *VisualizationBuilder
```

### <span class="badge object-method"></span> NullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```go
func (builder *VisualizationBuilder) NullValueMode(nullValueMode dashboardv2beta1.NullValueMode) *VisualizationBuilder
```

### <span class="badge object-method"></span> Override

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VisualizationBuilder) Override(systemRef string, matcher dashboardv2beta1.MatcherConfig, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByFieldType

Adds override rules for all the fields of the given type.

```go
func (builder *VisualizationBuilder) OverrideByFieldType(fieldType string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByName

Adds override rules for a specific field, referred to by its name.

```go
func (builder *VisualizationBuilder) OverrideByName(name string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByQuery

```go
func (builder *VisualizationBuilder) OverrideByQuery(queryRefId string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```go
func (builder *VisualizationBuilder) OverrideByRegexp(regexp string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> Overrides

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VisualizationBuilder) Overrides(overrides []cog.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```go
func (builder *VisualizationBuilder) Path(path string) *VisualizationBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *VisualizationBuilder) PointColor(pointColor string) *VisualizationBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *VisualizationBuilder) PointSize(pointSize float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> PointSymbol

```go
func (builder *VisualizationBuilder) PointSymbol(pointSymbol string) *VisualizationBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *VisualizationBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *VisualizationBuilder
```

### <span class="badge object-method"></span> ShowPoints

```go
func (builder *VisualizationBuilder) ShowPoints(showPoints common.VisibilityMode) *VisualizationBuilder
```

### <span class="badge object-method"></span> SpanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```go
func (builder *VisualizationBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Stacking

```go
func (builder *VisualizationBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Thresholds

Map numeric values to states

```go
func (builder *VisualizationBuilder) Thresholds(thresholds cog.Builder[dashboardv2beta1.ThresholdsConfig]) *VisualizationBuilder
```

### <span class="badge object-method"></span> ThresholdsStyle

```go
func (builder *VisualizationBuilder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *VisualizationBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Transform

```go
func (builder *VisualizationBuilder) Transform(transform common.GraphTransform) *VisualizationBuilder
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
func (builder *VisualizationBuilder) Unit(unit string) *VisualizationBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
