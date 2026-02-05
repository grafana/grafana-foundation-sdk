---
title: <span class="badge builder"></span> FieldColorBuilder
---
# <span class="badge builder"></span> FieldColorBuilder

## Constructor

```php
new FieldColorBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> fixedColor

The fixed color value for fixed or shades color modes.

```php
fixedColor(string $fixedColor)
```

### <span class="badge object-method"></span> mode

The main color scheme mode.

```php
mode(\Grafana\Foundation\Dashboard\FieldColorModeId $mode)
```

### <span class="badge object-method"></span> seriesBy

Some visualizations need to know how to assign a series color from by value color schemes.

```php
seriesBy(\Grafana\Foundation\Dashboard\FieldColorSeriesByMode $seriesBy)
```

## See also

 * <span class="badge object-type-class"></span> [FieldColor](./object-FieldColor.md)
