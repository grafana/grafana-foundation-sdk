---
title: <span class="badge object-type-class"></span> RepeatOptions
---
# <span class="badge object-type-class"></span> RepeatOptions

## Definition

```php
class RepeatOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\RepeatMode $mode;

    public string $value;

    public ?\Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection $direction;

    public ?int $maxPerRow;

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

 * <span class="badge builder"></span> [RepeatOptionsBuilder](./builder-RepeatOptionsBuilder.md)
