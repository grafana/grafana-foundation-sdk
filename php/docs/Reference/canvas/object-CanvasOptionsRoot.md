---
title: <span class="badge object-type-class"></span> CanvasOptionsRoot
---
# <span class="badge object-type-class"></span> CanvasOptionsRoot

## Definition

```php
class CanvasOptionsRoot implements \JsonSerializable
{
    /**
     * Name of the root element
     */
    public string $name;

    /**
     * Type of root element (frame)
     */
    public string $type;

    /**
     * The list of canvas elements attached to the root element
     * @var array<\Grafana\Foundation\Canvas\CanvasElementOptions>
     */
    public array $elements;

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

 * <span class="badge builder"></span> [CanvasOptionsRootBuilder](./builder-CanvasOptionsRootBuilder.md)
