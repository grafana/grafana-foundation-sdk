---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> geomap.Options
```

### <span class="badge object-method"></span> basemap

```python
def basemap(basemap: cogbuilder.Builder[common.MapLayerOptions]) -> typing.Self
```

### <span class="badge object-method"></span> controls

```python
def controls(controls: cogbuilder.Builder[geomap.ControlsOptions]) -> typing.Self
```

### <span class="badge object-method"></span> layers

```python
def layers(layers: list[cogbuilder.Builder[common.MapLayerOptions]]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

```python
def tooltip(tooltip: cogbuilder.Builder[geomap.TooltipOptions]) -> typing.Self
```

### <span class="badge object-method"></span> view

```python
def view(view: cogbuilder.Builder[geomap.MapViewConfig]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
