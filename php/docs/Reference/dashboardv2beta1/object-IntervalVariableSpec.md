---
title: <span class="badge object-type-class"></span> IntervalVariableSpec
---
# <span class="badge object-type-class"></span> IntervalVariableSpec

Interval variable specification

## Definition

```php
class IntervalVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $query;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>
     */
    public array $options;

    public bool $auto;

    public string $autoMin;

    public int $autoCount;

    public \Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

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

