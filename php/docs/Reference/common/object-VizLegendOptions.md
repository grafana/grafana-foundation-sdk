---
title: <span class="badge object-type-class"></span> VizLegendOptions
---
# <span class="badge object-type-class"></span> VizLegendOptions

TODO docs

## Definition

```php
class VizLegendOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\LegendDisplayMode $displayMode;

    public \Grafana\Foundation\Common\LegendPlacement $placement;

    public bool $showLegend;

    public ?bool $asTable;

    public ?bool $isVisible;

    public ?string $sortBy;

    public ?bool $sortDesc;

    public ?float $width;

    /**
     * @var array<string>
     */
    public array $calcs;

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

 * <span class="badge builder"></span> [VizLegendOptionsBuilder](./builder-VizLegendOptionsBuilder.md)
