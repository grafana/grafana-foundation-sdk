---
title: <span class="badge object-type-class"></span> NodeOptions
---
# <span class="badge object-type-class"></span> NodeOptions

## Definition

```php
class NodeOptions implements \JsonSerializable
{
    /**
     * Unit for the main stat to override what ever is set in the data frame.
     */
    public ?string $mainStatUnit;

    /**
     * Unit for the secondary stat to override what ever is set in the data frame.
     */
    public ?string $secondaryStatUnit;

    /**
     * Define which fields are shown as part of the node arc (colored circle around the node).
     * @var array<\Grafana\Foundation\Nodegraph\ArcOption>|null
     */
    public ?array $arcs;

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

 * <span class="badge builder"></span> [NodeOptionsBuilder](./builder-NodeOptionsBuilder.md)
