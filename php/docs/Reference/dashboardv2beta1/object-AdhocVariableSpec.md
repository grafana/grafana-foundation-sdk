---
title: <span class="badge object-type-class"></span> AdhocVariableSpec
---
# <span class="badge object-type-class"></span> AdhocVariableSpec

Adhoc variable specification

## Definition

```php
class AdhocVariableSpec implements \JsonSerializable
{
    public string $name;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>
     */
    public array $baseFilters;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>
     */
    public array $filters;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\MetricFindValue>
     */
    public array $defaultKeys;

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

