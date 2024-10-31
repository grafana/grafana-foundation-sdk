---
title: <span class="badge builder"></span> GraphFieldConfigBuilder
---
# <span class="badge builder"></span> GraphFieldConfigBuilder

## Constructor

```php
new GraphFieldConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
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

### <span class="badge object-method"></span> barAlignment

```php
barAlignment(\Grafana\Foundation\Common\BarAlignment $barAlignment)
```

### <span class="badge object-method"></span> barMaxWidth

```php
barMaxWidth(float $barMaxWidth)
```

### <span class="badge object-method"></span> barWidthFactor

```php
barWidthFactor(float $barWidthFactor)
```

### <span class="badge object-method"></span> drawStyle

```php
drawStyle(\Grafana\Foundation\Common\GraphDrawStyle $drawStyle)
```

### <span class="badge object-method"></span> fillBelowTo

```php
fillBelowTo(string $fillBelowTo)
```

### <span class="badge object-method"></span> fillColor

```php
fillColor(string $fillColor)
```

### <span class="badge object-method"></span> fillOpacity

```php
fillOpacity(float $fillOpacity)
```

### <span class="badge object-method"></span> gradientMode

```php
gradientMode(\Grafana\Foundation\Common\GraphGradientMode $gradientMode)
```

### <span class="badge object-method"></span> hideFrom

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom

```php
hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom)
```

### <span class="badge object-method"></span> insertNulls

@param bool|int $insertNulls

```php
insertNulls($insertNulls)
```

### <span class="badge object-method"></span> lineColor

```php
lineColor(string $lineColor)
```

### <span class="badge object-method"></span> lineInterpolation

```php
lineInterpolation(\Grafana\Foundation\Common\LineInterpolation $lineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineStyle> $lineStyle

```php
lineStyle(\Grafana\Foundation\Cog\Builder $lineStyle)
```

### <span class="badge object-method"></span> lineWidth

```php
lineWidth(float $lineWidth)
```

### <span class="badge object-method"></span> pointColor

```php
pointColor(string $pointColor)
```

### <span class="badge object-method"></span> pointSize

```php
pointSize(float $pointSize)
```

### <span class="badge object-method"></span> pointSymbol

```php
pointSymbol(string $pointSymbol)
```

### <span class="badge object-method"></span> scaleDistribution

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution

```php
scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution)
```

### <span class="badge object-method"></span> showPoints

```php
showPoints(\Grafana\Foundation\Common\VisibilityMode $showPoints)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

@param bool|float $spanNulls

```php
spanNulls($spanNulls)
```

### <span class="badge object-method"></span> stacking

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\StackingConfig> $stacking

```php
stacking(\Grafana\Foundation\Cog\Builder $stacking)
```

### <span class="badge object-method"></span> thresholdsStyle

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig> $thresholdsStyle

```php
thresholdsStyle(\Grafana\Foundation\Cog\Builder $thresholdsStyle)
```

### <span class="badge object-method"></span> transform

```php
transform(\Grafana\Foundation\Common\GraphTransform $transform)
```

## See also

 * <span class="badge object-type-class"></span> [GraphFieldConfig](./object-GraphFieldConfig.md)
