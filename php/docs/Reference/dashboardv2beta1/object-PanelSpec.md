---
title: <span class="badge object-type-class"></span> PanelSpec
---
# <span class="badge object-type-class"></span> PanelSpec

## Definition

```php
class PanelSpec implements \JsonSerializable
{
    public float $id;

    public string $title;

    public string $description;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\DataLink>
     */
    public array $links;

    public \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind $data;

    public \Grafana\Foundation\Dashboardv2beta1\VizConfigKind $vizConfig;

    public ?bool $transparent;

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

