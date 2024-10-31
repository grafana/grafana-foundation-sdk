---
title: <span class="badge builder"></span> DashboardRangeMapOptions
---
# <span class="badge builder"></span> DashboardRangeMapOptions

## Constructor

```python
DashboardRangeMapOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.DashboardRangeMapOptions
```

### <span class="badge object-method"></span> from_val

Min value of the range. It can be null which means -Infinity

```python
def from_val(from_val: float) -> typing.Self
```

### <span class="badge object-method"></span> result

Config to apply when the value is within the range

```python
def result(result: dashboard.ValueMappingResult) -> typing.Self
```

### <span class="badge object-method"></span> to

Max value of the range. It can be null which means +Infinity

```python
def to(to: float) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [DashboardRangeMapOptions](./object-DashboardRangeMapOptions.md)
