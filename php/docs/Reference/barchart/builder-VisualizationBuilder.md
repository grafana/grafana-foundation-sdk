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

### <span class="badge object-method"></span> barRadius

Controls the radius of each bar.

```php
barRadius(float $barRadius)
```

### <span class="badge object-method"></span> barWidth

Controls the width of bars. 1 = Max width, 0 = Min width.

```php
barWidth(float $barWidth)
```

### <span class="badge object-method"></span> colorByField

Use the color value for a sibling field to color each bar value.

```php
colorByField(string $colorByField)
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

Controls the fill opacity of the bars.

```php
fillOpacity(int $fillOpacity)
```

### <span class="badge object-method"></span> fullHighlight

Enables mode which highlights the entire bar area and shows tooltip when cursor

hovers over highlighted area

```php
fullHighlight(bool $fullHighlight)
```

### <span class="badge object-method"></span> gradientMode

Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.

Gradient appearance is influenced by the Fill opacity setting.

```php
gradientMode(\Grafana\Foundation\Common\GraphGradientMode $gradientMode)
```

### <span class="badge object-method"></span> groupWidth

Controls the width of groups. 1 = max with, 0 = min width.

```php
groupWidth(float $groupWidth)
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

Controls line width of the bars.

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

### <span class="badge object-method"></span> orientation

Controls the orientation of the bar chart, either vertical or horizontal.

```php
orientation(\Grafana\Foundation\Common\VizOrientation $orientation)
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

### <span class="badge object-method"></span> scaleDistribution

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution

```php
scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution)
```

### <span class="badge object-method"></span> showValue

This controls whether values are shown on top or to the left of bars.

```php
showValue(\Grafana\Foundation\Common\VisibilityMode $showValue)
```

### <span class="badge object-method"></span> stacking

Controls whether bars are stacked or not, either normally or in percent mode.

```php
stacking(\Grafana\Foundation\Common\StackingMode $stacking)
```

### <span class="badge object-method"></span> text

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text

```php
text(\Grafana\Foundation\Cog\Builder $text)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig> $thresholds

```php
thresholds(\Grafana\Foundation\Cog\Builder $thresholds)
```

### <span class="badge object-method"></span> thresholdsStyle

Threshold rendering

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig> $thresholdsStyle

```php
thresholdsStyle(\Grafana\Foundation\Cog\Builder $thresholdsStyle)
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

### <span class="badge object-method"></span> xField

Manually select which field from the dataset to represent the x field.

```php
xField(string $xField)
```

### <span class="badge object-method"></span> xTickLabelMaxLength

Sets the max length that a label can have before it is truncated.

```php
xTickLabelMaxLength(int $xTickLabelMaxLength)
```

### <span class="badge object-method"></span> xTickLabelRotation

Controls the rotation of the x axis labels.

```php
xTickLabelRotation(int $xTickLabelRotation)
```

### <span class="badge object-method"></span> xTickLabelSpacing

Controls the spacing between x axis labels.

negative values indicate backwards skipping behavior

```php
xTickLabelSpacing(int $xTickLabelSpacing)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
