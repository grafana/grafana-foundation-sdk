---
title: <span class="badge object-type-class"></span> TableFieldOptions
---
# <span class="badge object-type-class"></span> TableFieldOptions

Field options for each field within a table (e.g 10, "The String", 64.20, etc.)

Generally defines alignment, filtering capabilties, display options, etc.

## Definition

```php
class TableFieldOptions implements \JsonSerializable
{
    public ?float $width;

    public ?float $minWidth;

    public \Grafana\Foundation\Common\FieldTextAlignment $align;

    /**
     * This field is deprecated in favor of using cellOptions
     */
    public ?\Grafana\Foundation\Common\TableCellDisplayMode $displayMode;

    /**
     * @var \Grafana\Foundation\Common\TableAutoCellOptions|\Grafana\Foundation\Common\TableSparklineCellOptions|\Grafana\Foundation\Common\TableBarGaugeCellOptions|\Grafana\Foundation\Common\TableColoredBackgroundCellOptions|\Grafana\Foundation\Common\TableColorTextCellOptions|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions
     */
    public $cellOptions;

    /**
     * ?? default is missing or false ??
     */
    public ?bool $hidden;

    public bool $inspect;

    public ?bool $filterable;

    /**
     * Hides any header for a column, useful for columns that show some static content or buttons.
     */
    public ?bool $hideHeader;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [TableFieldOptionsBuilder](./builder-TableFieldOptionsBuilder.md)
