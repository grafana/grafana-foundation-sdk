---
title: <span class="badge builder"></span> MapLayerOptionsBuilder
---
# <span class="badge builder"></span> MapLayerOptionsBuilder

## Constructor

```php
new MapLayerOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> config

Custom options depending on the type

@param mixed $config

```php
config($config)
```

### <span class="badge object-method"></span> filterData

Defines a frame MatcherConfig that may filter data for the given layer

@param mixed $filterData

```php
filterData($filterData)
```

### <span class="badge object-method"></span> location

Common method to define geometry fields

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\FrameGeometrySource> $location

```php
location(\Grafana\Foundation\Cog\Builder $location)
```

### <span class="badge object-method"></span> name

configured unique display name

```php
name(string $name)
```

### <span class="badge object-method"></span> opacity

Common properties:

https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html

Layer opacity (0-1)

```php
opacity(int $opacity)
```

### <span class="badge object-method"></span> tooltip

Check tooltip (defaults to true)

```php
tooltip(bool $tooltip)
```

### <span class="badge object-method"></span> type

```php
type(string $type)
```

## See also

 * <span class="badge object-type-class"></span> [MapLayerOptions](./object-MapLayerOptions.md)
