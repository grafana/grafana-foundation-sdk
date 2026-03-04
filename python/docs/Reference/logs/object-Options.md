---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    show_labels: bool
    show_common_labels: bool
    show_time: bool
    show_log_context_toggle: bool
    wrap_log_message: bool
    prettify_log_message: bool
    enable_log_details: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy
    enable_infinite_scrolling: typing.Optional[bool]
    # Display controls to jump to the last or first log line, and filters by log level.
    show_controls: typing.Optional[bool]
    # Show a component to manage the displayed fields from the logs.
    show_field_selector: typing.Optional[bool]
    # Use a predefined coloring scheme to highlight relevant parts of the log lines.
    syntax_highlighting: typing.Optional[bool]
    font_size: typing.Optional[typing.Literal["default", "small"]]
    details_mode: typing.Optional[typing.Literal["inline", "sidebar"]]
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

