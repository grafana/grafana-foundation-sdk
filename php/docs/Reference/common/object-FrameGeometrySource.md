---
title: <span class="badge object-type-class"></span> FrameGeometrySource
---
# <span class="badge object-type-class"></span> FrameGeometrySource

## Definition

```php
class FrameGeometrySource implements \JsonSerializable
{
    public \Grafana\Foundation\Common\FrameGeometrySourceMode $mode;

    /**
     * Field mappings
     */
    public ?string $geohash;

    public ?string $latitude;

    public ?string $longitude;

    public ?string $wkt;

    public ?string $lookup;

    /**
     * Path to Gazetteer
     */
    public ?string $gazetteer;

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

 * <span class="badge builder"></span> [FrameGeometrySourceBuilder](./builder-FrameGeometrySourceBuilder.md)
