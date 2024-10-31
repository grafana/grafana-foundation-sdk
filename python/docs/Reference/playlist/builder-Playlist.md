---
title: <span class="badge builder"></span> Playlist
---
# <span class="badge builder"></span> Playlist

## Constructor

```python
Playlist()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> playlist.Playlist
```

### <span class="badge object-method"></span> interval

Interval sets the time between switching views in a playlist.

FIXME: Is this based on a standardized format or what options are available? Can datemath be used?

```python
def interval(interval: str) -> typing.Self
```

### <span class="badge object-method"></span> items

The ordered list of items that the playlist will iterate over.

FIXME! This should not be optional, but changing it makes the godegen awkward

```python
def items(items: list[cogbuilder.Builder[playlist.PlaylistItem]]) -> typing.Self
```

### <span class="badge object-method"></span> name

Name of the playlist.

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> uid

Unique playlist identifier. Generated on creation, either by the

creator of the playlist of by the application.

```python
def uid(uid: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Playlist](./object-Playlist.md)
