---
title: <span class="badge object-type-class"></span> FieldConfig
---
# <span class="badge object-type-class"></span> FieldConfig

## Definition

```php
class FieldConfig implements \JsonSerializable
{
    public ?int $lineWidth;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?int $fillOpacity;

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
