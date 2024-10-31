---
title: <span class="badge object-type-class"></span> CSVWave
---
# <span class="badge object-type-class"></span> CSVWave

## Definition

```php
class CSVWave implements \JsonSerializable
{
    public ?int $timeStep;

    public ?string $name;

    public ?string $valuesCSV;

    public ?string $labels;

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

 * <span class="badge builder"></span> [CSVWaveBuilder](./builder-CSVWaveBuilder.md)
