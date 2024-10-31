---
title: <span class="badge builder"></span> TableFieldOptionsBuilder
---
# <span class="badge builder"></span> TableFieldOptionsBuilder

## Constructor

```php
new TableFieldOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> align

```php
align(\Grafana\Foundation\Common\FieldTextAlignment $align)
```

### <span class="badge object-method"></span> cellOptions

@param \Grafana\Foundation\Common\TableAutoCellOptions|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableSparklineCellOptions>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableBarGaugeCellOptions>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableColoredBackgroundCellOptions>|\Grafana\Foundation\Common\TableColorTextCellOptions|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableDataLinksCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions $cellOptions

```php
cellOptions($cellOptions)
```

### <span class="badge object-method"></span> displayMode

This field is deprecated in favor of using cellOptions

```php
displayMode(\Grafana\Foundation\Common\TableCellDisplayMode $displayMode)
```

### <span class="badge object-method"></span> filterable

```php
filterable(bool $filterable)
```

### <span class="badge object-method"></span> hidden

?? default is missing or false ??

```php
hidden(bool $hidden)
```

### <span class="badge object-method"></span> hideHeader

Hides any header for a column, useful for columns that show some static content or buttons.

```php
hideHeader(bool $hideHeader)
```

### <span class="badge object-method"></span> inspect

```php
inspect(bool $inspect)
```

### <span class="badge object-method"></span> minWidth

```php
minWidth(float $minWidth)
```

### <span class="badge object-method"></span> width

```php
width(float $width)
```

## See also

 * <span class="badge object-type-class"></span> [TableFieldOptions](./object-TableFieldOptions.md)
