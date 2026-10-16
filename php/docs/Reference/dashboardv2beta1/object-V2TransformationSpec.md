---
title: <span class="badge object-type-class"></span> V2TransformationSpec
---
# <span class="badge object-type-class"></span> V2TransformationSpec

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Definition

```php
class V2TransformationSpec implements \JsonSerializable
{
    /**
     * Disabled transformations are skipped
     */
    public ?bool $disabled;

    /**
     * Optional frame matcher. When missing it will be applied to all results
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\MatcherConfig $filter;

    /**
     * Where to pull DataFrames from as input to transformation
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\DataTopic $topic;

    /**
     * Options to be passed to the transformer
     * Valid options depend on the transformer id
     * @var mixed
     */
    public $options;

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

 * <span class="badge builder"></span> [V2TransformationSpecBuilder](./builder-V2TransformationSpecBuilder.md)
