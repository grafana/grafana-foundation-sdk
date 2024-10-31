---
title: <span class="badge object-type-class"></span> MapViewConfig
---
# <span class="badge object-type-class"></span> MapViewConfig

## Definition

```python
class MapViewConfig:
    id_val: str
    lat: typing.Optional[int]
    lon: typing.Optional[int]
    zoom: typing.Optional[int]
    min_zoom: typing.Optional[int]
    max_zoom: typing.Optional[int]
    padding: typing.Optional[int]
    all_layers: typing.Optional[bool]
    last_only: typing.Optional[bool]
    layer: typing.Optional[str]
    shared: typing.Optional[bool]
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

 * <span class="badge builder"></span> [MapViewConfig](./builder-MapViewConfig.md)
