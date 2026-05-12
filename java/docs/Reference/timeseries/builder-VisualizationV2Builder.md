---
title: <span class="badge builder"></span> VisualizationV2Builder
---
# <span class="badge builder"></span> VisualizationV2Builder

## Constructor

```java
new VisualizationV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public VisualizationV2 build()
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

```java
public VisualizationV2Builder actions(List<com.grafana.foundation.cog.Builder<Action>> actions)
```

### <span class="badge object-method"></span> axisBorderShow

```java
public VisualizationV2Builder axisBorderShow(Boolean axisBorderShow)
```

### <span class="badge object-method"></span> axisCenteredZero

```java
public VisualizationV2Builder axisCenteredZero(Boolean axisCenteredZero)
```

### <span class="badge object-method"></span> axisColorMode

```java
public VisualizationV2Builder axisColorMode(AxisColorMode axisColorMode)
```

### <span class="badge object-method"></span> axisGridShow

```java
public VisualizationV2Builder axisGridShow(Boolean axisGridShow)
```

### <span class="badge object-method"></span> axisLabel

```java
public VisualizationV2Builder axisLabel(String axisLabel)
```

### <span class="badge object-method"></span> axisPlacement

```java
public VisualizationV2Builder axisPlacement(AxisPlacement axisPlacement)
```

### <span class="badge object-method"></span> axisSoftMax

```java
public VisualizationV2Builder axisSoftMax(Double axisSoftMax)
```

### <span class="badge object-method"></span> axisSoftMin

```java
public VisualizationV2Builder axisSoftMin(Double axisSoftMin)
```

### <span class="badge object-method"></span> axisWidth

```java
public VisualizationV2Builder axisWidth(Double axisWidth)
```

### <span class="badge object-method"></span> barAlignment

```java
public VisualizationV2Builder barAlignment(BarAlignment barAlignment)
```

### <span class="badge object-method"></span> barMaxWidth

```java
public VisualizationV2Builder barMaxWidth(Double barMaxWidth)
```

### <span class="badge object-method"></span> barWidthFactor

```java
public VisualizationV2Builder barWidthFactor(Double barWidthFactor)
```

### <span class="badge object-method"></span> colorScheme

Panel color configuration

```java
public VisualizationV2Builder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color)
```

### <span class="badge object-method"></span> dataLinks

The behavior when clicking on a result

```java
public VisualizationV2Builder dataLinks(List<Object> links)
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```java
public VisualizationV2Builder decimals(Double decimals)
```

### <span class="badge object-method"></span> description

Human readable field metadata

```java
public VisualizationV2Builder description(String description)
```

### <span class="badge object-method"></span> displayName

The display value for this field.  This supports template variables blank is auto

```java
public VisualizationV2Builder displayName(String displayName)
```

### <span class="badge object-method"></span> displayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```java
public VisualizationV2Builder displayNameFromDS(String displayNameFromDS)
```

### <span class="badge object-method"></span> drawStyle

```java
public VisualizationV2Builder drawStyle(GraphDrawStyle drawStyle)
```

### <span class="badge object-method"></span> fieldMinMax

Calculate min max per field

```java
public VisualizationV2Builder fieldMinMax(Boolean fieldMinMax)
```

### <span class="badge object-method"></span> fillBelowTo

```java
public VisualizationV2Builder fillBelowTo(String fillBelowTo)
```

### <span class="badge object-method"></span> fillColor

```java
public VisualizationV2Builder fillColor(String fillColor)
```

### <span class="badge object-method"></span> fillOpacity

```java
public VisualizationV2Builder fillOpacity(Double fillOpacity)
```

### <span class="badge object-method"></span> gradientMode

```java
public VisualizationV2Builder gradientMode(GraphGradientMode gradientMode)
```

### <span class="badge object-method"></span> hideFrom

```java
public VisualizationV2Builder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom)
```

### <span class="badge object-method"></span> insertNulls

```java
public VisualizationV2Builder insertNulls(BoolOrFloat64 insertNulls)
```

### <span class="badge object-method"></span> legend

```java
public VisualizationV2Builder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> lineColor

```java
public VisualizationV2Builder lineColor(String lineColor)
```

### <span class="badge object-method"></span> lineInterpolation

```java
public VisualizationV2Builder lineInterpolation(LineInterpolation lineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

```java
public VisualizationV2Builder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle)
```

### <span class="badge object-method"></span> lineWidth

```java
public VisualizationV2Builder lineWidth(Double lineWidth)
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

```java
public VisualizationV2Builder mappings(List<ValueMapping> mappings)
```

### <span class="badge object-method"></span> max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VisualizationV2Builder max(Double max)
```

### <span class="badge object-method"></span> min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VisualizationV2Builder min(Double min)
```

### <span class="badge object-method"></span> noValue

Alternative to empty string

```java
public VisualizationV2Builder noValue(String noValue)
```

### <span class="badge object-method"></span> nullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```java
public VisualizationV2Builder nullValueMode(NullValueMode nullValueMode)
```

### <span class="badge object-method"></span> orientation

```java
public VisualizationV2Builder orientation(VizOrientation orientation)
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

```java
public VisualizationV2Builder override(MatcherConfig matcher, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByFieldType

Adds override rules for all the fields of the given type.

```java
public VisualizationV2Builder overrideByFieldType(String fieldType, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByName

Adds override rules for a specific field, referred to by its name.

```java
public VisualizationV2Builder overrideByName(String name, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByQuery

```java
public VisualizationV2Builder overrideByQuery(String queryRefId, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```java
public VisualizationV2Builder overrideByRegexp(String regexp, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

```java
public VisualizationV2Builder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2FieldConfigSourceOverrides>> overrides)
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```java
public VisualizationV2Builder path(String path)
```

### <span class="badge object-method"></span> pointColor

```java
public VisualizationV2Builder pointColor(String pointColor)
```

### <span class="badge object-method"></span> pointSize

```java
public VisualizationV2Builder pointSize(Double pointSize)
```

### <span class="badge object-method"></span> pointSymbol

```java
public VisualizationV2Builder pointSymbol(String pointSymbol)
```

### <span class="badge object-method"></span> scaleDistribution

```java
public VisualizationV2Builder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution)
```

### <span class="badge object-method"></span> showPoints

```java
public VisualizationV2Builder showPoints(VisibilityMode showPoints)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```java
public VisualizationV2Builder spanNulls(BoolOrFloat64 spanNulls)
```

### <span class="badge object-method"></span> stacking

```java
public VisualizationV2Builder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```java
public VisualizationV2Builder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds)
```

### <span class="badge object-method"></span> thresholdsStyle

```java
public VisualizationV2Builder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle)
```

### <span class="badge object-method"></span> timezone

```java
public VisualizationV2Builder timezone(List<String> timezone)
```

### <span class="badge object-method"></span> tooltip

```java
public VisualizationV2Builder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip)
```

### <span class="badge object-method"></span> transform

```java
public VisualizationV2Builder transform(GraphTransform transform)
```

### <span class="badge object-method"></span> unit

Unit a field should use. The unit you select is applied to all fields except time.

You can use the units ID available in Grafana or a custom unit.

Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts

As custom unit, you can use the following formats:

`suffix:<suffix>` for custom unit that should go after value.

`prefix:<prefix>` for custom unit that should go before value.

`time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.

`si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.

`count:<unit>` for a custom count unit.

`currency:<unit>` for custom a currency unit.

```java
public VisualizationV2Builder unit(String unit)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.VizConfigKind](../dashboardv2/object-VizConfigKind.md)
