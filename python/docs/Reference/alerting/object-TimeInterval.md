---
title: <span class="badge object-type-class"></span> TimeInterval
---
# <span class="badge object-type-class"></span> TimeInterval

## Definition

```python
class TimeInterval:
    times: typing.Optional[list[alerting.TimeRange]]
    weekdays: typing.Optional[list[alerting.WeekdayRange]]
    days_of_month: typing.Optional[list[alerting.DayOfMonthRange]]
    months: typing.Optional[list[alerting.MonthRange]]
    years: typing.Optional[list[alerting.YearRange]]
    location: typing.Optional[alerting.Location]
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

 * <span class="badge builder"></span> [TimeInterval](./builder-TimeInterval.md)
