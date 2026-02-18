---
title: <span class="badge builder"></span> Item
---
# <span class="badge builder"></span> Item

## Constructor

```python
Item()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> playlistv0alpha1.Item
```

### <span class="badge object-method"></span> type_val

type of the item.

```python
def type_val(type_val: typing.Literal["dashboard_by_tag", "dashboard_by_uid", "dashboard_by_id"]) -> typing.Self
```

### <span class="badge object-method"></span> value

Value depends on type and describes the playlist item.

 - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This

 is not portable as the numerical identifier is non-deterministic between different instances.

 Will be replaced by dashboard_by_uid in the future. (deprecated)

 - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All

 dashboards behind the tag will be added to the playlist.

 - dashboard_by_uid: The value is the dashboard UID

```python
def value(value: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Item](./object-Item.md)
