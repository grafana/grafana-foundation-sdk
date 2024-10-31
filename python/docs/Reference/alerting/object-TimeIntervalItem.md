---
title: <span class="badge object-type-class"></span> TimeIntervalItem
---
# <span class="badge object-type-class"></span> TimeIntervalItem

## Definition

```python
class TimeIntervalItem:
    days_of_month: typing.Optional[list[str]]
    location: typing.Optional[str]
    months: typing.Optional[list[str]]
    times: typing.Optional[list[alerting.TimeIntervalTimeRange]]
    weekdays: typing.Optional[list[str]]
    years: typing.Optional[list[str]]
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

 * <span class="badge builder"></span> [TimeIntervalItem](./builder-TimeIntervalItem.md)
