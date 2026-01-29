---
title: <span class="badge object-type-class"></span> VariableValueOption
---
# <span class="badge object-type-class"></span> VariableValueOption

FIXME: should we introduce this? --- Variable value option

## Definition

```php
class VariableValueOption implements \JsonSerializable
{
    public string $label;

    /**
     * @var string|bool|float|\Grafana\Foundation\Dashboardv2beta1\CustomVariableValue
     */
    public $value;

    public ?string $group;

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

 * <span class="badge builder"></span> [VariableValueOptionBuilder](./builder-VariableValueOptionBuilder.md)
