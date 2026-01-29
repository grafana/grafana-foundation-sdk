---
title: <span class="badge object-type-class"></span> QueryGroupSpec
---
# <span class="badge object-type-class"></span> QueryGroupSpec

## Definition

```php
class QueryGroupSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>
     */
    public array $queries;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\TransformationKind>
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

