---
title: <span class="badge object-type-class"></span> Dashboardv2beta1SpecialValueMapOptions
---
# <span class="badge object-type-class"></span> Dashboardv2beta1SpecialValueMapOptions

## Definition

```php
class Dashboardv2beta1SpecialValueMapOptions implements \JsonSerializable
{
    /**
     * Special value to match against
     */
    public \Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch $match;

    /**
     * Config to apply when the value matches the special value
     */
    public \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result;

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

 * <span class="badge builder"></span> [Dashboardv2beta1SpecialValueMapOptionsBuilder](./builder-Dashboardv2beta1SpecialValueMapOptionsBuilder.md)
