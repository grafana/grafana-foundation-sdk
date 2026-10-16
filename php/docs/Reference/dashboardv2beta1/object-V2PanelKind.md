---
title: <span class="badge object-type-class"></span> V2PanelKind
---
# <span class="badge object-type-class"></span> V2PanelKind

The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,

which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has

to be forked rather than shared because PanelKind reaches TransformationKind through

QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not

change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.

## Definition

```php
class V2PanelKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec $spec;

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

 * <span class="badge builder"></span> [V2PanelKindBuilder](./builder-V2PanelKindBuilder.md)
