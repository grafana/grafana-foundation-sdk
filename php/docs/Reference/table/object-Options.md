---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    /**
     * Represents the index of the selected frame
     */
    public float $frameIndex;

    /**
     * Controls whether the panel should show the header
     */
    public bool $showHeader;

    /**
     * Controls whether the header should show icons for the column types
     */
    public ?bool $showTypeIcons;

    /**
     * Used to control row sorting
     * @var array<\Grafana\Foundation\Common\TableSortByFieldState>|null
     */
    public ?array $sortBy;

    /**
     * Controls footer options
     */
    public ?\Grafana\Foundation\Common\TableFooterOptions $footer;

    /**
     * Controls the height of the rows
     */
    public ?\Grafana\Foundation\Common\TableCellHeight $cellHeight;

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

