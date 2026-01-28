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

    public ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions $repeat;

    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind
     */
    public $layout;

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

 * <span class="badge builder"></span> [RowsLayoutRowSpecBuilder](./builder-RowsLayoutRowSpecBuilder.md)
