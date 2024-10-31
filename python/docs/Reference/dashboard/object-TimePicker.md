---
title: <span class="badge object-type-class"></span> TimePicker
---
# <span class="badge object-type-class"></span> TimePicker

## Definition

```python
class TimePicker:
    # Whether timepicker is visible or not.
    hidden: bool
    # Interval options available in the refresh picker dropdown.
    refresh_intervals: list[str]
    # Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
    collapse: bool
    # Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
    enable: bool
    # Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    time_options: list[str]
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

 * <span class="badge builder"></span> [TimePicker](./builder-TimePicker.md)
