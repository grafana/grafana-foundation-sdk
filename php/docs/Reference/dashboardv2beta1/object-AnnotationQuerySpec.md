---
title: <span class="badge object-type-class"></span> AnnotationQuerySpec
---
# <span class="badge object-type-class"></span> AnnotationQuerySpec

## Definition

```php
class AnnotationQuerySpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query;

    public bool $enable;

    public bool $hide;

    public string $iconColor;

    public string $name;

    public ?bool $builtIn;

    public ?\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter $filter;

    /**
     * Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
     */
    public ?string $placement;

    /**
     * Catch-all field for datasource-specific properties. Should not be available in as code tooling.
     * @var array<string, mixed>|null
     */
    public ?array $legacyOptions;

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

