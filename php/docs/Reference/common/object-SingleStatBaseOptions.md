---
title: <span class="badge object-type-class"></span> SingleStatBaseOptions
---
# <span class="badge object-type-class"></span> SingleStatBaseOptions

TODO docs

## Definition

```php
class SingleStatBaseOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

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

 * <span class="badge builder"></span> [SingleStatBaseOptionsBuilder](./builder-SingleStatBaseOptionsBuilder.md)
