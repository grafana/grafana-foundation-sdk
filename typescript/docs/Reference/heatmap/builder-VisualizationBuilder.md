---
title: <span class="badge builder"></span> VisualizationBuilder
---
# <span class="badge builder"></span> VisualizationBuilder

## Constructor

```typescript
new VisualizationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

```typescript
actions(actions: cog.Builder<dashboardv2beta1.Action>[])
```

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```typescript
calculate(calculate: boolean)
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

```typescript
calculation(calculation: cog.Builder<common.HeatmapCalculationOptions>)
```

### <span class="badge object-method"></span> cellGap

Controls gap between cells

```typescript
cellGap(cellGap: number)
```

### <span class="badge object-method"></span> cellRadius

Controls cell radius

```typescript
cellRadius(cellRadius: number)
```

### <span class="badge object-method"></span> cellValues

Controls cell value unit

```typescript
cellValues(cellValues: cog.Builder<heatmap.CellValues>)
```

### <span class="badge object-method"></span> color

Controls the color options

```typescript
color(color: cog.Builder<heatmap.HeatmapColorOptions>)
```

### <span class="badge object-method"></span> colorScheme

Panel color configuration

```typescript
colorScheme(color: cog.Builder<dashboardv2beta1.FieldColor>)
```

### <span class="badge object-method"></span> dataLinks

The behavior when clicking on a result

```typescript
dataLinks(links: any[])
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```typescript
decimals(decimals: number)
```

### <span class="badge object-method"></span> description

Human readable field metadata

```typescript
description(description: string)
```

### <span class="badge object-method"></span> displayName

The display value for this field.  This supports template variables blank is auto

```typescript
displayName(displayName: string)
```

### <span class="badge object-method"></span> displayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```typescript
displayNameFromDS(displayNameFromDS: string)
```

### <span class="badge object-method"></span> exemplars

Controls exemplar options

```typescript
exemplars(exemplars: cog.Builder<heatmap.ExemplarConfig>)
```

### <span class="badge object-method"></span> fieldMinMax

Calculate min max per field

```typescript
fieldMinMax(fieldMinMax: boolean)
```

### <span class="badge object-method"></span> filterValues

Filters values between a given range

```typescript
filterValues(filterValues: cog.Builder<heatmap.FilterValueRange>)
```

### <span class="badge object-method"></span> hideFrom

```typescript
hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>)
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```typescript
legend(legend: cog.Builder<heatmap.HeatmapLegend>)
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

```typescript
mappings(mappings: dashboardv2beta1.ValueMapping[])
```

### <span class="badge object-method"></span> max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```typescript
max(max: number)
```

### <span class="badge object-method"></span> min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```typescript
min(min: number)
```

### <span class="badge object-method"></span> noValue

Alternative to empty string

```typescript
noValue(noValue: string)
```

### <span class="badge object-method"></span> nullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```typescript
nullValueMode(nullValueMode: dashboardv2beta1.NullValueMode)
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

```typescript
override(override: {
	// Describes config override rules created when interacting with Grafana.
	__systemRef?: string;
	matcher: dashboardv2beta1.MatcherConfig;
	properties: dashboardv2beta1.DynamicConfigValue[];
})
```

### <span class="badge object-method"></span> overrideByFieldType

Adds override rules for all the fields of the given type.

```typescript
overrideByFieldType(fieldType: string, properties: dashboardv2beta1.DynamicConfigValue[])
```

### <span class="badge object-method"></span> overrideByName

Adds override rules for a specific field, referred to by its name.

```typescript
overrideByName(name: string, properties: dashboardv2beta1.DynamicConfigValue[])
```

### <span class="badge object-method"></span> overrideByQuery

```typescript
overrideByQuery(queryRefId: string, properties: dashboardv2beta1.DynamicConfigValue[])
```

### <span class="badge object-method"></span> overrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```typescript
overrideByRegexp(regexp: string, properties: dashboardv2beta1.DynamicConfigValue[])
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

```typescript
overrides(overrides: {
	// Describes config override rules created when interacting with Grafana.
	__systemRef?: string;
	matcher: dashboardv2beta1.MatcherConfig;
	properties: dashboardv2beta1.DynamicConfigValue[];
}[])
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```typescript
path(path: string)
```

### <span class="badge object-method"></span> rowsFrame

Controls tick alignment and value name when not calculating from data

```typescript
rowsFrame(rowsFrame: cog.Builder<heatmap.RowsHeatmapOptions>)
```

### <span class="badge object-method"></span> scaleDistribution

```typescript
scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>)
```

### <span class="badge object-method"></span> selectionMode

Controls which axis to allow selection on

```typescript
selectionMode(selectionMode: heatmap.HeatmapSelectionMode)
```

### <span class="badge object-method"></span> showValue

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```typescript
showValue(showValue: common.VisibilityMode)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```typescript
thresholds(thresholds: cog.Builder<dashboardv2beta1.ThresholdsConfig>)
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

```typescript
tooltip(tooltip: cog.Builder<heatmap.HeatmapTooltip>)
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

```typescript
unit(unit: string)
```

### <span class="badge object-method"></span> yAxis

Controls yAxis placement

```typescript
yAxis(yAxis: cog.Builder<heatmap.YAxisConfig>)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
