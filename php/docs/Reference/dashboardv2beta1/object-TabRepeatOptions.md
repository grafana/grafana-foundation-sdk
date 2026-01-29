---
title: <span class="badge object-type-class"></span> TabRepeatOptions
---
# <span class="badge object-type-class"></span> TabRepeatOptions

## Definition

```php
class TabRepeatOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\RepeatMode $mode;

    public string $value;

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

 * <span class="badge builder"></span> [TabRepeatOptionsBuilder](./builder-TabRepeatOptionsBuilder.md)
