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
    # TODO: figure out how to define callbacks
    on_click_filter_label: typing.Optional[object]
    on_click_filter_out_label: typing.Optional[object]
    is_filter_label_active: typing.Optional[object]
    on_click_filter_string: typing.Optional[object]
    on_click_filter_out_string: typing.Optional[object]
    on_click_show_field: typing.Optional[object]
    on_click_hide_field: typing.Optional[object]
    displayed_fields: typing.Optional[list[str]]
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

