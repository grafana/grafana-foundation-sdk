---
title: <span class="badge object-type-class"></span> RowsLayoutRowSpec
---
# <span class="badge object-type-class"></span> RowsLayoutRowSpec

## Definition

```php
class RowsLayoutRowSpec implements \JsonSerializable
{
    public ?string $title;

    public ?bool $collapse;

    public ?bool $hideHeader;

    public ?bool $fillScreen;

    public ?\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2\RowRepeatOptions $repeat;

    /**
     * @var \Grafana\Foundation\Dashboardv2\GridLayoutKind|\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\TabsLayoutKind|\Grafana\Foundation\Dashboardv2\RowsLayoutKind
     */
    public $layout;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\QueryVariableKind|\Grafana\Foundation\Dashboardv2\TextVariableKind|\Grafana\Foundation\Dashboardv2\ConstantVariableKind|\Grafana\Foundation\Dashboardv2\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2\IntervalVariableKind|\Grafana\Foundation\Dashboardv2\CustomVariableKind|\Grafana\Foundation\Dashboardv2\GroupByVariableKind|\Grafana\Foundation\Dashboardv2\AdhocVariableKind|\Grafana\Foundation\Dashboardv2\SwitchVariableKind>|null
     */
    public ?array $variables;

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

