---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Geomap\MapViewConfig $view;

    public \Grafana\Foundation\Geomap\ControlsOptions $controls;

    public \Grafana\Foundation\Common\MapLayerOptions $basemap;

    /**
     * @var array<\Grafana\Foundation\Common\MapLayerOptions>
     */
    public array $layers;

    public \Grafana\Foundation\Geomap\TooltipOptions $tooltip;

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

