---
title: <span class="badge object-type-class"></span> DashboardDashboardTemplating
---
# <span class="badge object-type-class"></span> DashboardDashboardTemplating

## Definition

```python
class DashboardDashboardTemplating:
    # List of configured template variables with their saved values along with some other metadata
    list_val: typing.Optional[list[dashboard.VariableModel]]
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

 * <span class="badge builder"></span> [DashboardDashboardTemplating](./builder-DashboardDashboardTemplating.md)
