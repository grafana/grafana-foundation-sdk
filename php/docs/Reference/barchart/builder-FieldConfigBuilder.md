---
title: <span class="badge builder"></span> FieldConfigBuilder
---
# <span class="badge builder"></span> FieldConfigBuilder

## Constructor

```php
new FieldConfigBuilder()
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

### <span class="badge object-method"></span> fillOpacity

Controls the fill opacity of the bars.

```php
fillOpacity(int $fillOpacity)
```

### <span class="badge object-method"></span> gradientMode

Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.

Gradient appearance is influenced by the Fill opacity setting.

```php
gradientMode(\Grafana\Foundation\Common\GraphGradientMode $gradientMode)
```

### <span class="badge object-method"></span> hideFrom

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom

```php
hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom)
```

### <span class="badge object-method"></span> lineWidth

Controls line width of the bars.

```php
lineWidth(int $lineWidth)
```

### <span class="badge object-method"></span> scaleDistribution

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution

```php
scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution)
```

### <span class="badge object-method"></span> thresholdsStyle

Threshold rendering

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig> $thresholdsStyle

```php
thresholdsStyle(\Grafana\Foundation\Cog\Builder $thresholdsStyle)
```

## See also

 * <span class="badge object-type-class"></span> [FieldConfig](./object-FieldConfig.md)
