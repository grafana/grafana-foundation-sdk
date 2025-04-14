---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    show_controls: bool
    show_time: bool
    wrap_log_message: bool
    enable_log_details: bool
    syntax_highlighting: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy
    grammar: typing.Optional[object]
    enable_infinite_scrolling: typing.Optional[bool]
    on_log_options_change: typing.Optional[object]
    on_new_logs_received: typing.Optional[object]
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

