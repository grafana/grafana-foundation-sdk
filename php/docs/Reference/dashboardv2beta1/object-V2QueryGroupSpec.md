---
title: <span class="badge object-type-class"></span> V2QueryGroupSpec
---
# <span class="badge object-type-class"></span> V2QueryGroupSpec

## Definition

```php
class V2QueryGroupSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>
     */
    public array $queries;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\V2TransformationKind>
     */
    public array $transformations;

    public \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec $queryOptions;

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

 * <span class="badge builder"></span> [V2QueryGroupSpecBuilder](./builder-V2QueryGroupSpecBuilder.md)
