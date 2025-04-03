---
title: <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder
---
# <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder

## Constructor

```php
new HeatmapCalculationBucketConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> mode

Sets the bucket calculation mode

```php
mode(\Grafana\Foundation\Common\HeatmapCalculationMode $mode)
```

### <span class="badge object-method"></span> scale

Controls the scale of the buckets

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scale

```php
scale(\Grafana\Foundation\Cog\Builder $scale)
```

### <span class="badge object-method"></span> value

The number of buckets to use for the axis in the heatmap

```php
value(string $value)
```

## See also

 * <span class="badge object-type-class"></span> [HeatmapCalculationBucketConfig](./object-HeatmapCalculationBucketConfig.md)
