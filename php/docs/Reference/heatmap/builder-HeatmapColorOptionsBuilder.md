---
title: <span class="badge builder"></span> HeatmapColorOptionsBuilder
---
# <span class="badge builder"></span> HeatmapColorOptionsBuilder

## Constructor

```php
new HeatmapColorOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> exponent

Controls the exponent when scale is set to exponential

```php
exponent(float $exponent)
```

### <span class="badge object-method"></span> fill

Controls the color fill when in opacity mode

```php
fill(string $fill)
```

### <span class="badge object-method"></span> max

Sets the maximum value for the color scale

```php
max(float $max)
```

### <span class="badge object-method"></span> min

Sets the minimum value for the color scale

```php
min(float $min)
```

### <span class="badge object-method"></span> mode

Sets the color mode

```php
mode(\Grafana\Foundation\Heatmap\HeatmapColorMode $mode)
```

### <span class="badge object-method"></span> reverse

Reverses the color scheme

```php
reverse(bool $reverse)
```

### <span class="badge object-method"></span> scale

Controls the color scale

```php
scale(\Grafana\Foundation\Heatmap\HeatmapColorScale $scale)
```

### <span class="badge object-method"></span> scheme

Controls the color scheme used

```php
scheme(string $scheme)
```

### <span class="badge object-method"></span> steps

Controls the number of color steps

```php
steps(int $steps)
```

## See also

 * <span class="badge object-type-class"></span> [HeatmapColorOptions](./object-HeatmapColorOptions.md)
