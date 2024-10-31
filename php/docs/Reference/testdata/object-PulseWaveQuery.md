---
title: <span class="badge object-type-class"></span> PulseWaveQuery
---
# <span class="badge object-type-class"></span> PulseWaveQuery

## Definition

```php
class PulseWaveQuery implements \JsonSerializable
{
    public ?int $offCount;

    public ?float $offValue;

    public ?int $onCount;

    public ?float $onValue;

    public ?int $timeStep;

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

 * <span class="badge builder"></span> [PulseWaveQueryBuilder](./builder-PulseWaveQueryBuilder.md)
