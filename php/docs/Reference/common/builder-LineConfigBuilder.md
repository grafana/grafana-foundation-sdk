---
title: <span class="badge builder"></span> LineConfigBuilder
---
# <span class="badge builder"></span> LineConfigBuilder

## Constructor

```php
new LineConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
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

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

@param bool|float $spanNulls

```php
spanNulls($spanNulls)
```

## See also

 * <span class="badge object-type-class"></span> [LineConfig](./object-LineConfig.md)
