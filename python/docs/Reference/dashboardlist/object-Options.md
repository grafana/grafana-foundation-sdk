---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    keep_time: bool
    include_vars: bool
    show_starred: bool
    show_recently_viewed: bool
    show_search: bool
    show_headings: bool
    max_items: int
    query: str
    folder_id: typing.Optional[int]
    tags: list[str]
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

