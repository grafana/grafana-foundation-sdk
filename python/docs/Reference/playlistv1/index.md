# playlistv1

## Objects

 * <span class="badge object-type-ref"></span> [Item](./object-Item.md)
 * <span class="badge object-type-class"></span> [Playlist](./object-Playlist.md)
 * <span class="badge object-type-scalar"></span> [PlaylistApiVersion](./object-PlaylistApiVersion.md)
 * <span class="badge object-type-class"></span> [PlaylistItem](./object-PlaylistItem.md)
 * <span class="badge object-type-scalar"></span> [PlaylistKind](./object-PlaylistKind.md)
## Builders

 * <span class="badge builder"></span> [Item](./builder-Item.md)
 * <span class="badge builder"></span> [Playlist](./builder-Playlist.md)
 * <span class="badge builder"></span> [PlaylistItem](./builder-PlaylistItem.md)
## Functions

### <span class="badge function"></span> manifest

Creates a resource manifest from a Playlist.

```python
def manifest(name: str, playlist: cogbuilder.Builder[playlistv1.Playlist]) -> Playlist
```

