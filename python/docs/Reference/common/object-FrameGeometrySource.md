---
title: <span class="badge object-type-class"></span> FrameGeometrySource
---
# <span class="badge object-type-class"></span> FrameGeometrySource

## Definition

```python
class FrameGeometrySource:
    mode: common.FrameGeometrySourceMode
    # Field mappings
    geohash: typing.Optional[str]
    latitude: typing.Optional[str]
    longitude: typing.Optional[str]
    wkt: typing.Optional[str]
    lookup: typing.Optional[str]
    # Path to Gazetteer
    gazetteer: typing.Optional[str]
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

 * <span class="badge builder"></span> [FrameGeometrySource](./builder-FrameGeometrySource.md)
