---
title: <span class="badge object-type-class"></span> Preferences
---
# <span class="badge object-type-class"></span> Preferences

Dashboard specific preferences (applied per dashboard = all users using the dashboard)

## Definition

```python
class Preferences:
    """
    Dashboard specific preferences (applied per dashboard = all users using the dashboard)
    """

    # default layout template to be used when new containers are created
    layout: typing.Optional[typing.Union[dashboardv2.AutoGridLayoutKind, dashboardv2.GridLayoutKind]]
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

 * <span class="badge builder"></span> [Preferences](./builder-Preferences.md)
