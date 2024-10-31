---
title: <span class="badge object-type-class"></span> DateHistogramSettings
---
# <span class="badge object-type-class"></span> DateHistogramSettings

## Definition

```php
class DateHistogramSettings implements \JsonSerializable
{
    public ?string $interval;

    public ?string $minDocCount;

    public ?string $trimEdges;

    public ?string $offset;

    public ?string $timeZone;

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

 * <span class="badge builder"></span> [DateHistogramSettingsBuilder](./builder-DateHistogramSettingsBuilder.md)
