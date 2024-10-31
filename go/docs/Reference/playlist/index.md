# playlist

## Objects

 * <span class="badge object-type-struct"></span> [Playlist](./object-Playlist.md)
 * <span class="badge object-type-struct"></span> [PlaylistItem](./object-PlaylistItem.md)
 * <span class="badge object-type-enum"></span> [PlaylistItemType](./object-PlaylistItemType.md)
## Builders

 * <span class="badge builder"></span> [PlaylistBuilder](./builder-PlaylistBuilder.md)
 * <span class="badge builder"></span> [PlaylistItemBuilder](./builder-PlaylistItemBuilder.md)
## Functions

### <span class="badge function"></span> PlaylistConverter

PlaylistConverter accepts a `Playlist` object and generates the Go code to build this object using builders.

```go
func PlaylistConverter(input Playlist) string
```

### <span class="badge function"></span> PlaylistItemConverter

PlaylistItemConverter accepts a `PlaylistItem` object and generates the Go code to build this object using builders.

```go
func PlaylistItemConverter(input PlaylistItem) string
```

