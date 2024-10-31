---
title: <span class="badge object-type-class"></span> StreamingQuery
---
# <span class="badge object-type-class"></span> StreamingQuery

## Definition

```php
class StreamingQuery implements \JsonSerializable
{
    public \Grafana\Foundation\Testdata\StreamingQueryType $type;

    public int $speed;

    public int $spread;

    public int $noise;

    public ?int $bands;

    public ?string $url;

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

 * <span class="badge builder"></span> [StreamingQueryBuilder](./builder-StreamingQueryBuilder.md)
