# playlistv0alpha1

## Objects

 * <span class="badge object-type-struct"></span> [Item](./object-Item.md)
 * <span class="badge object-type-enum"></span> [ItemType](./object-ItemType.md)
 * <span class="badge object-type-struct"></span> [Playlist](./object-Playlist.md)
 * <span class="badge object-type-scalar"></span> [PlaylistKind](./object-PlaylistKind.md)
 * <span class="badge object-type-scalar"></span> [PlaylistV0Alpha1](./object-PlaylistV0Alpha1.md)
## Builders

 * <span class="badge builder"></span> [ItemBuilder](./builder-ItemBuilder.md)
 * <span class="badge builder"></span> [PlaylistBuilder](./builder-PlaylistBuilder.md)
## Functions

### <span class="badge function"></span> NewItem

NewItem creates a new Item object.

```go
func NewItem() *Item
```

### <span class="badge function"></span> NewPlaylist

NewPlaylist creates a new Playlist object.

```go
func NewPlaylist() *Playlist
```

### <span class="badge function"></span> Manifest

Creates a resource manifest from a Playlist.

```go
func Manifest(name string, playlist cog.Builder[Playlist]) *PlaylistBuilder
```

### <span class="badge function"></span> ItemConverter

ItemConverter accepts a `Item` object and generates the Go code to build this object using builders.

```go
func ItemConverter(input Item) string
```

### <span class="badge function"></span> PlaylistConverter

PlaylistConverter accepts a `Playlist` object and generates the Go code to build this object using builders.

```go
func PlaylistConverter(input Playlist) string
```

