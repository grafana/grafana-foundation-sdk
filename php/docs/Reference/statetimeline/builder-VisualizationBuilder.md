---
title: <span class="badge builder"></span> VisualizationBuilder
---
# <span class="badge builder"></span> VisualizationBuilder

## Constructor

```php
new VisualizationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Action>> $actions

```php
actions(array $actions)
```

### <span class="badge object-method"></span> alignValue

Controls value alignment on the timelines

```php
alignValue(\Grafana\Foundation\Common\TimelineValueAlignment $alignValue)
```

### <span class="badge object-method"></span> axisBorderShow

```php
axisBorderShow(bool $axisBorderShow)
```

### <span class="badge object-method"></span> axisCenteredZero

```php
axisCenteredZero(bool $axisCenteredZero)
```

### <span class="badge object-method"></span> axisColorMode

```php
axisColorMode(\Grafana\Foundation\Common\AxisColorMode $axisColorMode)
```

### <span class="badge object-method"></span> axisGridShow

```php
axisGridShow(bool $axisGridShow)
```

### <span class="badge object-method"></span> axisLabel

```php
axisLabel(string $axisLabel)
```

### <span class="badge object-method"></span> axisPlacement

```php
axisPlacement(\Grafana\Foundation\Common\AxisPlacement $axisPlacement)
```

### <span class="badge object-method"></span> axisSoftMax

```php
axisSoftMax(float $axisSoftMax)
```

### <span class="badge object-method"></span> axisSoftMin

```php
axisSoftMin(float $axisSoftMin)
```

### <span class="badge object-method"></span> axisWidth

```php
axisWidth(float $axisWidth)
```

### <span class="badge object-method"></span> colorScheme

Panel color configuration

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\FieldColor> $color

```php
colorScheme(\Grafana\Foundation\Cog\Builder $color)
```

### <span class="badge object-method"></span> dataLinks

The behavior when clicking on a result

@param array<mixed> $links

```php
dataLinks(array $links)
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```php
decimals(float $decimals)
```

### <span class="badge object-method"></span> description

Human readable field metadata

```php
description(string $description)
```

### <span class="badge object-method"></span> displayName

The display value for this field.  This supports template variables blank is auto

```php
displayName(string $displayName)
```

### <span class="badge object-method"></span> displayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```php
displayNameFromDS(string $displayNameFromDS)
```

### <span class="badge object-method"></span> fieldMinMax

Calculate min max per field

```php
fieldMinMax(bool $fieldMinMax)
```

### <span class="badge object-method"></span> fillOpacity

```php
fillOpacity(int $fillOpacity)
```

### <span class="badge object-method"></span> hideFrom

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom

```php
hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> lineWidth

```php
lineWidth(int $lineWidth)
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

@param array<\Grafana\Foundation\Dashboardv2beta1\ValueMap|\Grafana\Foundation\Dashboardv2beta1\RangeMap|\Grafana\Foundation\Dashboardv2beta1\RegexMap|\Grafana\Foundation\Dashboardv2beta1\SpecialValueMap> $mappings

```php
mappings(array $mappings)
```

### <span class="badge object-method"></span> max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```php
max(float $max)
```

### <span class="badge object-method"></span> mergeValues

Merge equal consecutive values

```php
mergeValues(bool $mergeValues)
```

### <span class="badge object-method"></span> min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```php
min(float $min)
```

### <span class="badge object-method"></span> noValue

Alternative to empty string

```php
noValue(string $noValue)
```

### <span class="badge object-method"></span> nullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```php
nullValueMode(\Grafana\Foundation\Dashboardv2beta1\NullValueMode $nullValueMode)
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides> $override

```php
override(\Grafana\Foundation\Cog\Builder $override)
```

### <span class="badge object-method"></span> overrideByFieldType

Adds override rules for all the fields of the given type.

@param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties

```php
overrideByFieldType(string $fieldType, array $properties)
```

### <span class="badge object-method"></span> overrideByName

Adds override rules for a specific field, referred to by its name.

@param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties

```php
overrideByName(string $name, array $properties)
```

### <span class="badge object-method"></span> overrideByQuery

@param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties

```php
overrideByQuery(string $queryRefId, array $properties)
```

### <span class="badge object-method"></span> overrideByRegexp

Adds override rules for the fields whose name match the given regexp.

@param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties

```php
overrideByRegexp(string $regexp, array $properties)
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides>> $overrides

```php
overrides(array $overrides)
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```php
path(string $path)
```

### <span class="badge object-method"></span> perPage

Enables pagination when > 0

```php
perPage(float $perPage)
```

### <span class="badge object-method"></span> rowHeight

Controls the row height

```php
rowHeight(float $rowHeight)
```

### <span class="badge object-method"></span> scaleDistribution

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution

```php
scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution)
```

### <span class="badge object-method"></span> showValue

Show timeline values on chart

```php
showValue(\Grafana\Foundation\Common\VisibilityMode $showValue)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig> $thresholds

```php
thresholds(\Grafana\Foundation\Cog\Builder $thresholds)
```

### <span class="badge object-method"></span> timezone

@param array<string> $timezone

```php
timezone(array $timezone)
```

### <span class="badge object-method"></span> tooltip

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip

```php
tooltip(\Grafana\Foundation\Cog\Builder $tooltip)
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

```php
unit(string $unit)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
