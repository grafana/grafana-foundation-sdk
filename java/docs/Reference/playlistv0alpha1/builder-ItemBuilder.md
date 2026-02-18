---
title: <span class="badge builder"></span> ItemBuilder
---
# <span class="badge builder"></span> ItemBuilder

## Constructor

```java
new ItemBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Item build()
```

### <span class="badge object-method"></span> type

type of the item.

```java
public ItemBuilder type(ItemType type)
```

### <span class="badge object-method"></span> value

Value depends on type and describes the playlist item.

 - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This

 is not portable as the numerical identifier is non-deterministic between different instances.

 Will be replaced by dashboard_by_uid in the future. (deprecated)

 - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All

 dashboards behind the tag will be added to the playlist.

 - dashboard_by_uid: The value is the dashboard UID

```java
public ItemBuilder value(String value)
```

## See also

 * <span class="badge object-type-class"></span> [Item](./object-Item.md)
