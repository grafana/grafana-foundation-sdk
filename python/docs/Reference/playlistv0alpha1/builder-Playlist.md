---
title: <span class="badge builder"></span> Playlist
---
# <span class="badge builder"></span> Playlist

## Constructor

```python
Playlist(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> playlistv0alpha1.Playlist
```

### <span class="badge object-method"></span> interval

```python
def interval(interval: str) -> typing.Self
```

### <span class="badge object-method"></span> item

```python
def item(item: cogbuilder.Builder[playlistv0alpha1.Item]) -> typing.Self
```

### <span class="badge object-method"></span> items

```python
def items(items: list[cogbuilder.Builder[playlistv0alpha1.Item]]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Playlist](./object-Playlist.md)
