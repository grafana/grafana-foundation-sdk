---
title: <span class="badge builder"></span> VisualizationBuilder
---
# <span class="badge builder"></span> VisualizationBuilder

## Constructor

```java
new VisualizationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Visualization build()
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

```java
public VisualizationBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions)
```

### <span class="badge object-method"></span> axisBorderShow

```java
public VisualizationBuilder axisBorderShow(Boolean axisBorderShow)
```

### <span class="badge object-method"></span> axisCenteredZero

```java
public VisualizationBuilder axisCenteredZero(Boolean axisCenteredZero)
```

### <span class="badge object-method"></span> axisColorMode

```java
public VisualizationBuilder axisColorMode(AxisColorMode axisColorMode)
```

### <span class="badge object-method"></span> axisGridShow

```java
public VisualizationBuilder axisGridShow(Boolean axisGridShow)
```

### <span class="badge object-method"></span> axisLabel

```java
public VisualizationBuilder axisLabel(String axisLabel)
```

### <span class="badge object-method"></span> axisPlacement

```java
public VisualizationBuilder axisPlacement(AxisPlacement axisPlacement)
```

### <span class="badge object-method"></span> axisSoftMax

```java
public VisualizationBuilder axisSoftMax(Double axisSoftMax)
```

### <span class="badge object-method"></span> axisSoftMin

```java
public VisualizationBuilder axisSoftMin(Double axisSoftMin)
```

### <span class="badge object-method"></span> axisWidth

```java
public VisualizationBuilder axisWidth(Double axisWidth)
```

### <span class="badge object-method"></span> barAlignment

```java
public VisualizationBuilder barAlignment(BarAlignment barAlignment)
```

### <span class="badge object-method"></span> barMaxWidth

```java
public VisualizationBuilder barMaxWidth(Double barMaxWidth)
```

### <span class="badge object-method"></span> barWidthFactor

```java
public VisualizationBuilder barWidthFactor(Double barWidthFactor)
```

### <span class="badge object-method"></span> colorScheme

Panel color configuration

```java
public VisualizationBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color)
```

### <span class="badge object-method"></span> dataLinks

The behavior when clicking on a result

```java
public VisualizationBuilder dataLinks(List<Object> links)
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```java
public VisualizationBuilder decimals(Double decimals)
```

### <span class="badge object-method"></span> description

Human readable field metadata

```java
public VisualizationBuilder description(String description)
```

### <span class="badge object-method"></span> displayName

The display value for this field.  This supports template variables blank is auto

```java
public VisualizationBuilder displayName(String displayName)
```

### <span class="badge object-method"></span> displayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```java
public VisualizationBuilder displayNameFromDS(String displayNameFromDS)
```

### <span class="badge object-method"></span> drawStyle

```java
public VisualizationBuilder drawStyle(GraphDrawStyle drawStyle)
```

### <span class="badge object-method"></span> fillBelowTo

```java
public VisualizationBuilder fillBelowTo(String fillBelowTo)
```

### <span class="badge object-method"></span> fillColor

```java
public VisualizationBuilder fillColor(String fillColor)
```

### <span class="badge object-method"></span> fillOpacity

```java
public VisualizationBuilder fillOpacity(Double fillOpacity)
```

### <span class="badge object-method"></span> gradientMode

```java
public VisualizationBuilder gradientMode(GraphGradientMode gradientMode)
```

### <span class="badge object-method"></span> hideFrom

```java
public VisualizationBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom)
```

### <span class="badge object-method"></span> insertNulls

```java
public VisualizationBuilder insertNulls(BoolOrFloat64 insertNulls)
```

### <span class="badge object-method"></span> legend

```java
public VisualizationBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> lineColor

```java
public VisualizationBuilder lineColor(String lineColor)
```

### <span class="badge object-method"></span> lineInterpolation

```java
public VisualizationBuilder lineInterpolation(LineInterpolation lineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

```java
public VisualizationBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle)
```

### <span class="badge object-method"></span> lineWidth

```java
public VisualizationBuilder lineWidth(Double lineWidth)
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

```java
public VisualizationBuilder mappings(List<ValueMapping> mappings)
```

### <span class="badge object-method"></span> max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VisualizationBuilder max(Double max)
```

### <span class="badge object-method"></span> min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VisualizationBuilder min(Double min)
```

### <span class="badge object-method"></span> noValue

Alternative to empty string

```java
public VisualizationBuilder noValue(String noValue)
```

### <span class="badge object-method"></span> orientation

```java
public VisualizationBuilder orientation(VizOrientation orientation)
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

```java
public VisualizationBuilder override(MatcherConfig matcher, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByFieldType

Adds override rules for all the fields of the given type.

```java
public VisualizationBuilder overrideByFieldType(String fieldType, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByName

Adds override rules for a specific field, referred to by its name.

```java
public VisualizationBuilder overrideByName(String name, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByQuery

```java
public VisualizationBuilder overrideByQuery(String queryRefId, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```java
public VisualizationBuilder overrideByRegexp(String regexp, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

```java
public VisualizationBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides)
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```java
public VisualizationBuilder path(String path)
```

### <span class="badge object-method"></span> pointColor

```java
public VisualizationBuilder pointColor(String pointColor)
```

### <span class="badge object-method"></span> pointSize

```java
public VisualizationBuilder pointSize(Double pointSize)
```

### <span class="badge object-method"></span> pointSymbol

```java
public VisualizationBuilder pointSymbol(String pointSymbol)
```

### <span class="badge object-method"></span> scaleDistribution

```java
public VisualizationBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution)
```

### <span class="badge object-method"></span> showPoints

```java
public VisualizationBuilder showPoints(VisibilityMode showPoints)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```java
public VisualizationBuilder spanNulls(BoolOrFloat64 spanNulls)
```

### <span class="badge object-method"></span> stacking

```java
public VisualizationBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```java
public VisualizationBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds)
```

### <span class="badge object-method"></span> thresholdsStyle

```java
public VisualizationBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle)
```

### <span class="badge object-method"></span> timezone

```java
public VisualizationBuilder timezone(List<String> timezone)
```

### <span class="badge object-method"></span> tooltip

```java
public VisualizationBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip)
```

### <span class="badge object-method"></span> transform

```java
public VisualizationBuilder transform(GraphTransform transform)
```

### <span class="badge object-method"></span> unit

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

```java
public VisualizationBuilder unit(String unit)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
