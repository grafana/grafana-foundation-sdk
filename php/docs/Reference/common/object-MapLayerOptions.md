---
title: <span class="badge object-type-class"></span> MapLayerOptions
---
# <span class="badge object-type-class"></span> MapLayerOptions

## Definition

```php
class MapLayerOptions implements \JsonSerializable
{
    public string $type;

    /**
     * configured unique display name
     */
    public string $name;

    /**
     * Custom options depending on the type
     * @var mixed|null
     */
    public $config;

    /**
     * Common method to define geometry fields
     */
    public ?\Grafana\Foundation\Common\FrameGeometrySource $location;

    /**
     * Defines a frame MatcherConfig that may filter data for the given layer
     * @var mixed|null
     */
    public $filterData;

    /**
     * Common properties:
     * https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
     * Layer opacity (0-1)
     */
    public ?int $opacity;

    /**
     * Check tooltip (defaults to true)
     */
    public ?bool $tooltip;

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

 * <span class="badge builder"></span> [MapLayerOptionsBuilder](./builder-MapLayerOptionsBuilder.md)
