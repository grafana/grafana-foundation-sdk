---
title: <span class="badge object-type-class"></span> PublicDashboard
---
# <span class="badge object-type-class"></span> PublicDashboard

## Definition

```python
class PublicDashboard:
    # Unique public dashboard identifier
    uid: str
    # Dashboard unique identifier referenced by this public dashboard
    dashboard_uid: str
    # Unique public access token
    access_token: typing.Optional[str]
    # Flag that indicates if the public dashboard is enabled
    is_enabled: bool
    # Flag that indicates if annotations are enabled
    annotations_enabled: bool
    # Flag that indicates if the time range picker is enabled
    time_selection_enabled: bool
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

 * <span class="badge builder"></span> [PublicDashboard](./builder-PublicDashboard.md)
