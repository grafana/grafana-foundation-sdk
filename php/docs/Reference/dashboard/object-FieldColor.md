---
title: <span class="badge object-type-class"></span> FieldColor
---
# <span class="badge object-type-class"></span> FieldColor

Map a field to a color.

## Definition

```php
class FieldColor implements \JsonSerializable
{
    /**
     * The main color scheme mode.
     */
    public \Grafana\Foundation\Dashboard\FieldColorModeId $mode;

    /**
     * The fixed color value for fixed or shades color modes.
     */
    public ?string $fixedColor;

    /**
     * Some visualizations need to know how to assign a series color from by value color schemes.
     */
    public ?\Grafana\Foundation\Dashboard\FieldColorSeriesByMode $seriesBy;

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

 * <span class="badge builder"></span> [FieldColorBuilder](./builder-FieldColorBuilder.md)
