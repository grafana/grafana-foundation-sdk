---
title: <span class="badge object-type-class"></span> MapViewConfig
---
# <span class="badge object-type-class"></span> MapViewConfig

## Definition

```php
class MapViewConfig implements \JsonSerializable
{
    public string $id;

    public ?int $lat;

    public ?int $lon;

    public ?int $zoom;

    public ?int $minZoom;

    public ?int $maxZoom;

    public ?int $padding;

    public ?bool $allLayers;

    public ?bool $lastOnly;

    public ?string $layer;

    public ?bool $shared;

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

 * <span class="badge builder"></span> [MapViewConfigBuilder](./builder-MapViewConfigBuilder.md)
