---
title: <span class="badge object-type-class"></span> VizConfigSpec
---
# <span class="badge object-type-class"></span> VizConfigSpec

--- Kinds ---

## Definition

```php
class VizConfigSpec implements \JsonSerializable
{
    /**
     * @var mixed|null
     */
    public $options;

    public \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource $fieldConfig;

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

