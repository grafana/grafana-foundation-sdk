---
title: <span class="badge builder"></span> ThresholdsConfigBuilder
---
# <span class="badge builder"></span> ThresholdsConfigBuilder

## Constructor

```php
new ThresholdsConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> mode

Thresholds mode.

```php
mode(\Grafana\Foundation\Dashboard\ThresholdsMode $mode)
```

### <span class="badge object-method"></span> steps

Must be sorted by 'value', first value is always -Infinity

@param array<\Grafana\Foundation\Dashboard\Threshold> $steps

```php
steps(array $steps)
```

## See also

 * <span class="badge object-type-class"></span> [ThresholdsConfig](./object-ThresholdsConfig.md)
