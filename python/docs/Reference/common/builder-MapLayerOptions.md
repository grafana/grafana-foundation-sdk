---
title: <span class="badge builder"></span> MapLayerOptions
---
# <span class="badge builder"></span> MapLayerOptions

## Constructor

```python
MapLayerOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> common.MapLayerOptions
```

### <span class="badge object-method"></span> config

Custom options depending on the type

```python
def config(config: object) -> typing.Self
```

### <span class="badge object-method"></span> filter_data

Defines a frame MatcherConfig that may filter data for the given layer

```python
def filter_data(filter_data: object) -> typing.Self
```

### <span class="badge object-method"></span> location

Common method to define geometry fields

```python
def location(location: cogbuilder.Builder[common.FrameGeometrySource]) -> typing.Self
```

### <span class="badge object-method"></span> name

configured unique display name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> opacity

Common properties:

https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html

Layer opacity (0-1)

```python
def opacity(opacity: int) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

Check tooltip (defaults to true)

```python
def tooltip(tooltip: bool) -> typing.Self
```

### <span class="badge object-method"></span> type_val

```python
def type_val(type_val: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [MapLayerOptions](./object-MapLayerOptions.md)
