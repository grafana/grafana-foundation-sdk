---
title: <span class="badge object-type-class"></span> GraphPanel
---
# <span class="badge object-type-class"></span> GraphPanel

Support for legacy graph panel.

@deprecated this a deprecated panel type

## Definition

```php
class GraphPanel implements \JsonSerializable
{
    public string $type;

    /**
     * @deprecated this is part of deprecated graph panel
     */
    public ?\Grafana\Foundation\Dashboard\DashboardGraphPanelLegend $legend;

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

 * <span class="badge builder"></span> [GraphPanelBuilder](./builder-GraphPanelBuilder.md)
