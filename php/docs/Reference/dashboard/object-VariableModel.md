---
title: <span class="badge object-type-class"></span> VariableModel
---
# <span class="badge object-type-class"></span> VariableModel

A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.

## Definition

```php
class VariableModel implements \JsonSerializable
{
    /**
     * Type of variable
     */
    public \Grafana\Foundation\Dashboard\VariableType $type;

    /**
     * Name of variable
     */
    public string $name;

    /**
     * Optional display name
     */
    public ?string $label;

    /**
     * Visibility configuration for the variable
     */
    public ?\Grafana\Foundation\Dashboard\VariableHide $hide;

    /**
     * Whether the variable value should be managed by URL query params or not
     */
    public ?bool $skipUrlSync;

    /**
     * Description of variable. It can be defined but `null`.
     */
    public ?string $description;

    /**
     * Query used to fetch values for a variable
     * @var string|array<string, mixed>|null
     */
    public $query;

    /**
     * Data source used to fetch values for a variable. It can be defined but `null`.
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Shows current selected variable text/value on the dashboard
     */
    public ?\Grafana\Foundation\Dashboard\VariableOption $current;

    /**
     * Whether multiple values can be selected or not from variable value list
     */
    public ?bool $multi;

    /**
     * Options that can be selected for a variable.
     * @var array<\Grafana\Foundation\Dashboard\VariableOption>|null
     */
    public ?array $options;

    public ?\Grafana\Foundation\Dashboard\VariableRefresh $refresh;

    /**
     * Options sort order
     */
    public ?\Grafana\Foundation\Dashboard\VariableSort $sort;

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

 * <span class="badge builder"></span> [AdHocVariableBuilder](./builder-AdHocVariableBuilder.md)
 * <span class="badge builder"></span> [ConstantVariableBuilder](./builder-ConstantVariableBuilder.md)
 * <span class="badge builder"></span> [CustomVariableBuilder](./builder-CustomVariableBuilder.md)
 * <span class="badge builder"></span> [DatasourceVariableBuilder](./builder-DatasourceVariableBuilder.md)
 * <span class="badge builder"></span> [IntervalVariableBuilder](./builder-IntervalVariableBuilder.md)
 * <span class="badge builder"></span> [QueryVariableBuilder](./builder-QueryVariableBuilder.md)
 * <span class="badge builder"></span> [TextBoxVariableBuilder](./builder-TextBoxVariableBuilder.md)
