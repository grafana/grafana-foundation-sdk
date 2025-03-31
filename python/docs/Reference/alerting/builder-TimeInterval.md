---
title: <span class="badge builder"></span> TimeInterval
---
# <span class="badge builder"></span> TimeInterval

## Constructor

```python
TimeInterval()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.TimeInterval
```

### <span class="badge object-method"></span> days_of_month

```python
def days_of_month(days_of_month: list[cogbuilder.Builder[alerting.DayOfMonthRange]]) -> typing.Self
```

### <span class="badge object-method"></span> location

```python
def location(location: alerting.Location) -> typing.Self
```

### <span class="badge object-method"></span> months

```python
def months(months: list[cogbuilder.Builder[alerting.MonthRange]]) -> typing.Self
```

### <span class="badge object-method"></span> times

```python
def times(times: list[cogbuilder.Builder[alerting.TimeRange]]) -> typing.Self
```

### <span class="badge object-method"></span> weekdays

```python
def weekdays(weekdays: list[cogbuilder.Builder[alerting.WeekdayRange]]) -> typing.Self
```

### <span class="badge object-method"></span> years

```python
def years(years: list[cogbuilder.Builder[alerting.YearRange]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TimeInterval](./object-TimeInterval.md)
