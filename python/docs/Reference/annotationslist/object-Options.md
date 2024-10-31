---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    only_from_this_dashboard: bool
    only_in_time_range: bool
    tags: list[str]
    limit: int
    show_user: bool
    show_time: bool
    show_tags: bool
    navigate_to_panel: bool
    navigate_before: str
    navigate_after: str
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

