---
title: <span class="badge object-type-class"></span> ExtendedStat
---
# <span class="badge object-type-class"></span> ExtendedStat

## Definition

```php
class ExtendedStat implements \JsonSerializable
{
    public string $label;

    public \Grafana\Foundation\Elasticsearch\ExtendedStatMetaType $value;

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

 * <span class="badge builder"></span> [ExtendedStatBuilder](./builder-ExtendedStatBuilder.md)