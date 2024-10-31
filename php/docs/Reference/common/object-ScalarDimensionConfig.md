---
title: <span class="badge object-type-class"></span> ScalarDimensionConfig
---
# <span class="badge object-type-class"></span> ScalarDimensionConfig

## Definition

```php
class ScalarDimensionConfig implements \JsonSerializable
{
    public float $min;

    public float $max;

    public ?float $fixed;

    /**
     * fixed: T -- will be added by each element
     */
    public ?string $field;

    public ?\Grafana\Foundation\Common\ScalarDimensionMode $mode;

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

 * <span class="badge builder"></span> [ScalarDimensionConfigBuilder](./builder-ScalarDimensionConfigBuilder.md)
