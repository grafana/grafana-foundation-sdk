---
title: <span class="badge object-type-class"></span> CanvasConnection
---
# <span class="badge object-type-class"></span> CanvasConnection

## Definition

```php
class CanvasConnection implements \JsonSerializable
{
    public \Grafana\Foundation\Canvas\ConnectionCoordinates $source;

    public \Grafana\Foundation\Canvas\ConnectionCoordinates $target;

    public ?string $targetName;

    public \Grafana\Foundation\Canvas\ConnectionPath $path;

    public ?\Grafana\Foundation\Common\ColorDimensionConfig $color;

    public ?\Grafana\Foundation\Common\ScaleDimensionConfig $size;

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

 * <span class="badge builder"></span> [CanvasConnectionBuilder](./builder-CanvasConnectionBuilder.md)
