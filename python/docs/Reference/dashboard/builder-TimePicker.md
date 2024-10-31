---
title: <span class="badge builder"></span> TimePicker
---
# <span class="badge builder"></span> TimePicker

## Constructor

```python
TimePicker()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.TimePicker
```

### <span class="badge object-method"></span> collapse

Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.

```python
def collapse(collapse: bool) -> typing.Self
```

### <span class="badge object-method"></span> enable

Whether timepicker is enabled or not. Has no effect on provisioned dashboard.

```python
def enable(enable: bool) -> typing.Self
```

### <span class="badge object-method"></span> hidden

Whether timepicker is visible or not.

```python
def hidden(hidden: bool) -> typing.Self
```

### <span class="badge object-method"></span> refresh_intervals

Interval options available in the refresh picker dropdown.

```python
def refresh_intervals(refresh_intervals: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> time_options

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

```python
def time_options(time_options: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TimePicker](./object-TimePicker.md)