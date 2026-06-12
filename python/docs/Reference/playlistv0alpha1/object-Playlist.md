---
title: <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Playlist
---
# <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Playlist

<span class="badge deprecated"></span>Prefer using playlistv1.Playlist instead.

## Definition

```python
class Playlist:
    warnings.warn("Prefer using playlistv1.Playlist instead.", DeprecationWarning)
    title: str
    interval: str
    items: list[playlistv0alpha1.Item]
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

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [Playlist](./builder-Playlist.md)
