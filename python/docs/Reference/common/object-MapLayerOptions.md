---
title: <span class="badge object-type-class"></span> MapLayerOptions
---
# <span class="badge object-type-class"></span> MapLayerOptions

## Definition

```python
class MapLayerOptions:
    type_val: str
    # configured unique display name
    name: str
    # Custom options depending on the type
    config: typing.Optional[object]
    # Common method to define geometry fields
    location: typing.Optional[common.FrameGeometrySource]
    # Defines a frame MatcherConfig that may filter data for the given layer
    filter_data: typing.Optional[object]
    # Common properties:
    # https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
    # Layer opacity (0-1)
    opacity: typing.Optional[int]
    # Check tooltip (defaults to true)
    tooltip: typing.Optional[bool]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [MapLayerOptions](./builder-MapLayerOptions.md)
