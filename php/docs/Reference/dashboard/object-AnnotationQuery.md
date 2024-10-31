---
title: <span class="badge object-type-class"></span> AnnotationQuery
---
# <span class="badge object-type-class"></span> AnnotationQuery

TODO docs

FROM: AnnotationQuery in grafana-data/src/types/annotations.ts

## Definition

```php
class AnnotationQuery implements \JsonSerializable
{
    /**
     * Name of annotation.
     */
    public string $name;

    /**
     * Datasource where the annotations data is
     */
    public \Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * When enabled the annotation query is issued with every dashboard refresh
     */
    public bool $enable;

    /**
     * Annotation queries can be toggled on or off at the top of the dashboard.
     * When hide is true, the toggle is not shown in the dashboard.
     */
    public ?bool $hide;

    /**
     * Color to use for the annotation event markers
     */
    public string $iconColor;

    /**
     * Filters to apply when fetching annotations
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationPanelFilter $filter;

    /**
     * TODO.. this should just be a normal query target
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationTarget $target;

    /**
     * TODO -- this should not exist here, it is based on the --grafana-- datasource
     */
    public ?string $type;

    public ?string $expr;

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

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
