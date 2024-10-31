---
title: <span class="badge object-type-class"></span> TimePickerConfig
---
# <span class="badge object-type-class"></span> TimePickerConfig

Time picker configuration

It defines the default config for the time picker and the refresh picker for the specific dashboard.

## Definition

```python
class TimePickerConfig:
    """
    Time picker configuration
    It defines the default config for the time picker and the refresh picker for the specific dashboard.
    """

    # Whether timepicker is visible or not.
    hidden: typing.Optional[bool]
    # Interval options available in the refresh picker dropdown.
    refresh_intervals: typing.Optional[list[str]]
    # Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    time_options: typing.Optional[list[str]]
    # Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    now_delay: typing.Optional[str]
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
