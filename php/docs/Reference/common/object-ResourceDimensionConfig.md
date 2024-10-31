---
title: <span class="badge object-type-class"></span> ResourceDimensionConfig
---
# <span class="badge object-type-class"></span> ResourceDimensionConfig

Links to a resource (image/svg path)

## Definition

```php
class ResourceDimensionConfig implements \JsonSerializable
{
    public \Grafana\Foundation\Common\ResourceDimensionMode $mode;

    /**
     * fixed: T -- will be added by each element
     */
    public ?string $field;

    public ?string $fixed;

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

 * <span class="badge builder"></span> [ResourceDimensionConfigBuilder](./builder-ResourceDimensionConfigBuilder.md)
