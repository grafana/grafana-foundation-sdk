---
title: <span class="badge builder"></span> ItemBuilder
---
# <span class="badge builder"></span> ItemBuilder

## Constructor

```go
func NewItemBuilder() *ItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ItemBuilder) Build() (Item, error)
```

### <span class="badge object-method"></span> Type

type of the item.

```go
func (builder *ItemBuilder) Type(typeArg playlistv0alpha1.ItemType) *ItemBuilder
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
func (builder *ItemBuilder) Value(value string) *ItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Item](./object-Item.md)
