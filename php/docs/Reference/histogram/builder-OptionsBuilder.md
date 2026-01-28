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

### <span class="badge object-method"></span> bucketCount

Bucket count (approx)

```php
bucketCount(int $bucketCount)
```

### <span class="badge object-method"></span> bucketOffset

Offset buckets by this amount

```php
bucketOffset(float $bucketOffset)
```

### <span class="badge object-method"></span> bucketSize

Size of each bucket

```php
bucketSize(int $bucketSize)
```

### <span class="badge object-method"></span> combine

Combines multiple series into a single histogram

```php
combine(bool $combine)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> tooltip

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip

```php
tooltip(\Grafana\Foundation\Cog\Builder $tooltip)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
