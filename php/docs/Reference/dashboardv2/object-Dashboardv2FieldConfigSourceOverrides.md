---
title: <span class="badge object-type-class"></span> Dashboardv2FieldConfigSourceOverrides
---
# <span class="badge object-type-class"></span> Dashboardv2FieldConfigSourceOverrides

## Definition

```php
class Dashboardv2FieldConfigSourceOverrides implements \JsonSerializable
{
    public ?string $systemRef;

    public \Grafana\Foundation\Dashboardv2\MatcherConfig $matcher;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\DynamicConfigValue>
     */
    public array $properties;

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

 * <span class="badge builder"></span> [Dashboardv2FieldConfigSourceOverridesBuilder](./builder-Dashboardv2FieldConfigSourceOverridesBuilder.md)
