---
title: <span class="badge object-type-class"></span> QueryVariableSpec
---
# <span class="badge object-type-class"></span> QueryVariableSpec

Query variable specification

## Definition

```php
class QueryVariableSpec implements \JsonSerializable
{
    public string $name;

    public \Grafana\Foundation\Dashboardv2\VariableOption $current;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2\VariableHide $hide;

    public \Grafana\Foundation\Dashboardv2\VariableRefresh $refresh;

    public bool $skipUrlSync;

    public ?string $description;

    public \Grafana\Foundation\Dashboardv2\DataQueryKind $query;

    public string $regex;

    public ?\Grafana\Foundation\Dashboardv2\VariableRegexApplyTo $regexApplyTo;

    public \Grafana\Foundation\Dashboardv2\VariableSort $sort;

    public ?string $definition;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\VariableOption>
     */
    public array $options;

    public bool $multi;

    public bool $includeAll;

    public ?string $allValue;

    public ?string $placeholder;

    public bool $allowCustomValue;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\VariableOption>|null
     */
    public ?array $staticOptions;

    public ?\Grafana\Foundation\Dashboardv2\QueryVariableSpecStaticOptionsOrder $staticOptionsOrder;

    public ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin;

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

