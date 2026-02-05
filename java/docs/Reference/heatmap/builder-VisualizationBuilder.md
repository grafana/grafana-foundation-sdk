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

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```java
public VisualizationBuilder calculate(Boolean calculate)
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

```java
public VisualizationBuilder calculation(com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> calculation)
```

### <span class="badge object-method"></span> cellGap

Controls gap between cells

```java
public VisualizationBuilder cellGap(Integer cellGap)
```

### <span class="badge object-method"></span> cellRadius

Controls cell radius

```java
public VisualizationBuilder cellRadius(Float cellRadius)
```

### <span class="badge object-method"></span> cellValues

Controls cell value unit

```java
public VisualizationBuilder cellValues(com.grafana.foundation.cog.Builder<CellValues> cellValues)
```

### <span class="badge object-method"></span> color

Controls the color options

```java
public VisualizationBuilder color(com.grafana.foundation.cog.Builder<HeatmapColorOptions> color)
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

### <span class="badge object-method"></span> exemplars

Controls exemplar options

```java
public VisualizationBuilder exemplars(com.grafana.foundation.cog.Builder<ExemplarConfig> exemplars)
```

### <span class="badge object-method"></span> filterValues

Filters values between a given range

```java
public VisualizationBuilder filterValues(com.grafana.foundation.cog.Builder<FilterValueRange> filterValues)
```

### <span class="badge object-method"></span> hideFrom

```java
public VisualizationBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom)
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```java
public VisualizationBuilder legend(com.grafana.foundation.cog.Builder<HeatmapLegend> legend)
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

### <span class="badge object-method"></span> rowsFrame

Controls tick alignment and value name when not calculating from data

```java
public VisualizationBuilder rowsFrame(com.grafana.foundation.cog.Builder<RowsHeatmapOptions> rowsFrame)
```

### <span class="badge object-method"></span> scaleDistribution

```java
public VisualizationBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution)
```

### <span class="badge object-method"></span> selectionMode

Controls which axis to allow selection on

```java
public VisualizationBuilder selectionMode(HeatmapSelectionMode selectionMode)
```

### <span class="badge object-method"></span> showValue

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```java
public VisualizationBuilder showValue(VisibilityMode showValue)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```java
public VisualizationBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds)
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

```java
public VisualizationBuilder tooltip(com.grafana.foundation.cog.Builder<HeatmapTooltip> tooltip)
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

### <span class="badge object-method"></span> yAxis

Controls yAxis placement

```java
public VisualizationBuilder yAxis(com.grafana.foundation.cog.Builder<YAxisConfig> yAxis)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
