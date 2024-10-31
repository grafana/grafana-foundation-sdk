---
title: <span class="badge object-type-class"></span> AnnotationTarget
---
# <span class="badge object-type-class"></span> AnnotationTarget

TODO: this should be a regular DataQuery that depends on the selected dashboard

these match the properties of the "grafana" datasouce that is default in most dashboards

## Definition

```php
class AnnotationTarget implements \JsonSerializable
{
    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public int $limit;

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public bool $matchAny;

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     * @var array<string>
     */
    public array $tags;

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public string $type;

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

 * <span class="badge builder"></span> [AnnotationTargetBuilder](./builder-AnnotationTargetBuilder.md)
