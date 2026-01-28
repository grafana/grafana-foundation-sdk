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

### <span class="badge object-method"></span> cellHeight

Controls the height of the rows

```php
cellHeight(\Grafana\Foundation\Common\TableCellHeight $cellHeight)
```

### <span class="badge object-method"></span> footer

Controls footer options

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableFooterOptions> $footer

```php
footer(\Grafana\Foundation\Cog\Builder $footer)
```

### <span class="badge object-method"></span> frameIndex

Represents the index of the selected frame

```php
frameIndex(float $frameIndex)
```

### <span class="badge object-method"></span> showHeader

Controls whether the panel should show the header

```php
showHeader(bool $showHeader)
```

### <span class="badge object-method"></span> showTypeIcons

Controls whether the header should show icons for the column types

```php
showTypeIcons(bool $showTypeIcons)
```

### <span class="badge object-method"></span> sortBy

Used to control row sorting

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableSortByFieldState>> $sortBy

```php
sortBy(array $sortBy)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
