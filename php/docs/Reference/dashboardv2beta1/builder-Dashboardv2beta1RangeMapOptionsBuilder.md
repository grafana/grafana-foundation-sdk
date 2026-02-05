---
title: <span class="badge builder"></span> Dashboardv2beta1RangeMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2beta1RangeMapOptionsBuilder

## Constructor

```php
new Dashboardv2beta1RangeMapOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> from

Min value of the range. It can be null which means -Infinity

```php
from(float $from)
```

### <span class="badge object-method"></span> result

Config to apply when the value is within the range

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ValueMappingResult> $result

```php
result(\Grafana\Foundation\Cog\Builder $result)
```

### <span class="badge object-method"></span> to

Max value of the range. It can be null which means +Infinity

```php
to(float $to)
```

## See also

 * <span class="badge object-type-class"></span> [Dashboardv2beta1RangeMapOptions](./object-Dashboardv2beta1RangeMapOptions.md)
