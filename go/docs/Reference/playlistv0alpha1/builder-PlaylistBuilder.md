---
title: <span class="badge builder"></span> PlaylistBuilder
---
# <span class="badge builder"></span> PlaylistBuilder

## Constructor

```go
func NewPlaylistBuilder(title string) *PlaylistBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PlaylistBuilder) Build() (Playlist, error)
```

### <span class="badge object-method"></span> Interval

```go
func (builder *PlaylistBuilder) Interval(interval string) *PlaylistBuilder
```

### <span class="badge object-method"></span> Item

```go
func (builder *PlaylistBuilder) Item(item cog.Builder[playlistv0alpha1.Item]) *PlaylistBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *PlaylistBuilder) Items(items []cog.Builder[playlistv0alpha1.Item]) *PlaylistBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *PlaylistBuilder) Title(title string) *PlaylistBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Playlist](./object-Playlist.md)
