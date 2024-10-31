---
title: <span class="badge builder"></span> PlaylistItemBuilder
---
# <span class="badge builder"></span> PlaylistItemBuilder

## Constructor

```go
func NewPlaylistItemBuilder() *PlaylistItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PlaylistItemBuilder) Build() (PlaylistItem, error)
```

### <span class="badge object-method"></span> Title

Title is an unused property -- it will be removed in the future

```go
func (builder *PlaylistItemBuilder) Title(title string) *PlaylistItemBuilder
```

### <span class="badge object-method"></span> Type

Type of the item.

```go
func (builder *PlaylistItemBuilder) Type(typeArg playlist.PlaylistItemType) *PlaylistItemBuilder
```

### <span class="badge object-method"></span> Value

Value depends on type and describes the playlist item.



 - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This

 is not portable as the numerical identifier is non-deterministic between different instances.

 Will be replaced by dashboard_by_uid in the future. (deprecated)

 - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All

 dashboards behind the tag will be added to the playlist.

 - dashboard_by_uid: The value is the dashboard UID

```go
func (builder *PlaylistItemBuilder) Value(value string) *PlaylistItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PlaylistItem](./object-PlaylistItem.md)
