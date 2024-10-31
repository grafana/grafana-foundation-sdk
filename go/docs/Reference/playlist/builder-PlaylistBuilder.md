---
title: <span class="badge builder"></span> PlaylistBuilder
---
# <span class="badge builder"></span> PlaylistBuilder

## Constructor

```go
func NewPlaylistBuilder() *PlaylistBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PlaylistBuilder) Build() (Playlist, error)
```

### <span class="badge object-method"></span> Interval

Interval sets the time between switching views in a playlist.

FIXME: Is this based on a standardized format or what options are available? Can datemath be used?

```go
func (builder *PlaylistBuilder) Interval(interval string) *PlaylistBuilder
```

### <span class="badge object-method"></span> Items

The ordered list of items that the playlist will iterate over.

FIXME! This should not be optional, but changing it makes the godegen awkward

```go
func (builder *PlaylistBuilder) Items(items []cog.Builder[playlist.PlaylistItem]) *PlaylistBuilder
```

### <span class="badge object-method"></span> Name

Name of the playlist.

```go
func (builder *PlaylistBuilder) Name(name string) *PlaylistBuilder
```

### <span class="badge object-method"></span> Uid

Unique playlist identifier. Generated on creation, either by the

creator of the playlist of by the application.

```go
func (builder *PlaylistBuilder) Uid(uid string) *PlaylistBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Playlist](./object-Playlist.md)
