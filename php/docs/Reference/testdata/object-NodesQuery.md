---
title: <span class="badge object-type-class"></span> NodesQuery
---
# <span class="badge object-type-class"></span> NodesQuery

## Definition

```php
class NodesQuery implements \JsonSerializable
{
    public ?int $count;

    public ?int $seed;

    /**
     * Possible enum values:
     *  - `"random"` 
     *  - `"random edges"` 
     *  - `"response_medium"` 
     *  - `"response_small"` 
     *  - `"feature_showcase"` 
     */
    public ?\Grafana\Foundation\Testdata\NodesQueryType $type;

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

 * <span class="badge builder"></span> [NodesQueryBuilder](./builder-NodesQueryBuilder.md)
