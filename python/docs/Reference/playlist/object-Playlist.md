---
title: <span class="badge object-type-class"></span> Playlist
---
# <span class="badge object-type-class"></span> Playlist

## Definition

```python
class Playlist:
    # Unique playlist identifier. Generated on creation, either by the
    # creator of the playlist of by the application.
    uid: str
    # Name of the playlist.
    name: str
    # Interval sets the time between switching views in a playlist.
    # FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
    interval: str
    # The ordered list of items that the playlist will iterate over.
    # FIXME! This should not be optional, but changing it makes the godegen awkward
    items: typing.Optional[list[playlist.PlaylistItem]]
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

 * <span class="badge builder"></span> [Playlist](./builder-Playlist.md)
