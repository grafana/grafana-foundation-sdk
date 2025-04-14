---
title: <span class="badge object-type-class"></span> SimulationQuery
---
# <span class="badge object-type-class"></span> SimulationQuery

## Definition

```php
class SimulationQuery implements \JsonSerializable
{
    /**
     * @var mixed|null
     */
    public $config;

    public \Grafana\Foundation\Testdata\Key $key;

    public ?bool $last;

    public ?bool $stream;

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

 * <span class="badge builder"></span> [SimulationQueryBuilder](./builder-SimulationQueryBuilder.md)
