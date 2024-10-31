---
title: <span class="badge builder"></span> PlaylistItemBuilder
---
# <span class="badge builder"></span> PlaylistItemBuilder

## Constructor

```typescript
new PlaylistItemBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> title

Title is an unused property -- it will be removed in the future

```typescript
title(title: string)
```

### <span class="badge object-method"></span> type

Type of the item.

```typescript
type(type: "dashboard_by_uid" | "dashboard_by_id" | "dashboard_by_tag")
```

### <span class="badge object-method"></span> value

Value depends on type and describes the playlist item.



 - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This

 is not portable as the numerical identifier is non-deterministic between different instances.

 Will be replaced by dashboard_by_uid in the future. (deprecated)

 - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All

 dashboards behind the tag will be added to the playlist.

 - dashboard_by_uid: The value is the dashboard UID

```typescript
value(value: string)
```

## See also

 * <span class="badge object-type-interface"></span> [PlaylistItem](./object-PlaylistItem.md)
