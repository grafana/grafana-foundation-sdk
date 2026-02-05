---
title: <span class="badge object-type-class"></span> CustomVariableSpec
---
# <span class="badge object-type-class"></span> CustomVariableSpec

Custom variable specification

## Definition

```php
class CustomVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $query;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>
     */
    public array $options;

    public bool $multi;

    public bool $includeAll;

    public ?string $allValue;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    public bool $allowCustomValue;

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

