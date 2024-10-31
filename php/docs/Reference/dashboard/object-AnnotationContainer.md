---
title: <span class="badge object-type-class"></span> AnnotationContainer
---
# <span class="badge object-type-class"></span> AnnotationContainer

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

## Definition

```php
class AnnotationContainer implements \JsonSerializable
{
    /**
     * List of annotations
     * @var array<\Grafana\Foundation\Dashboard\AnnotationQuery>|null
     */
    public ?array $list;

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

