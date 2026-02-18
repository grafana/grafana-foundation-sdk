---
title: <span class="badge object-type-class"></span> Item
---
# <span class="badge object-type-class"></span> Item

## Definition

```python
class Item:
    # type of the item.
    type_val: typing.Literal["dashboard_by_tag", "dashboard_by_uid", "dashboard_by_id"]
    # Value depends on type and describes the playlist item.
    #  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    #  is not portable as the numerical identifier is non-deterministic between different instances.
    #  Will be replaced by dashboard_by_uid in the future. (deprecated)
    #  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    #  dashboards behind the tag will be added to the playlist.
    #  - dashboard_by_uid: The value is the dashboard UID
    value: str
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [Item](./builder-Item.md)
