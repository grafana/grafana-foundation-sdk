---
title: <span class="badge builder"></span> TimePickerBuilder
---
# <span class="badge builder"></span> TimePickerBuilder

## Constructor

```java
new TimePickerBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TimePicker build()
```

### <span class="badge object-method"></span> hidden

Whether timepicker is visible or not.

```java
public TimePickerBuilder hidden(Boolean hidden)
```

### <span class="badge object-method"></span> nowDelay

Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.

```java
public TimePickerBuilder nowDelay(String nowDelay)
```

### <span class="badge object-method"></span> quickRanges

Quick ranges for time picker.

```java
public TimePickerBuilder quickRanges(List<com.grafana.foundation.cog.Builder<TimeOption>> quickRanges)
```

### <span class="badge object-method"></span> refreshIntervals

Interval options available in the refresh picker dropdown.

```java
public TimePickerBuilder refreshIntervals(List<String> refreshIntervals)
```

## See also

 * <span class="badge object-type-class"></span> [TimePickerConfig](./object-TimePickerConfig.md)
