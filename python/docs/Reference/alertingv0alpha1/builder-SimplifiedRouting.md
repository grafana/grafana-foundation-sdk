---
title: <span class="badge builder"></span> SimplifiedRouting
---
# <span class="badge builder"></span> SimplifiedRouting

## Constructor

```python
SimplifiedRouting()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv0alpha1.SimplifiedRouting
```

### <span class="badge object-method"></span> active_time_intervals

```python
def active_time_intervals(active_time_intervals: list[alertingv0alpha1.TimeIntervalRef]) -> typing.Self
```

### <span class="badge object-method"></span> group_by

```python
def group_by(group_by: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> group_interval

```python
def group_interval(group_interval: alertingv0alpha1.PromDuration) -> typing.Self
```

### <span class="badge object-method"></span> group_wait

```python
def group_wait(group_wait: alertingv0alpha1.PromDuration) -> typing.Self
```

### <span class="badge object-method"></span> mute_time_intervals

```python
def mute_time_intervals(mute_time_intervals: list[alertingv0alpha1.TimeIntervalRef]) -> typing.Self
```

### <span class="badge object-method"></span> receiver

```python
def receiver(receiver: str) -> typing.Self
```

### <span class="badge object-method"></span> repeat_interval

```python
def repeat_interval(repeat_interval: alertingv0alpha1.PromDuration) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SimplifiedRouting](./object-SimplifiedRouting.md)
