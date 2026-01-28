---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```php
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
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

### <span class="badge object-method"></span> fullHighlight

Enables mode which highlights the entire bar area and shows tooltip when cursor

hovers over highlighted area

```php
fullHighlight(bool $fullHighlight)
```

### <span class="badge object-method"></span> groupWidth

Controls the width of groups. 1 = max with, 0 = min width.

```php
groupWidth(float $groupWidth)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> orientation

Controls the orientation of the bar chart, either vertical or horizontal.

```php
orientation(\Grafana\Foundation\Common\VizOrientation $orientation)
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

### <span class="badge object-method"></span> tooltip

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip

```php
tooltip(\Grafana\Foundation\Cog\Builder $tooltip)
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

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
