---
title: <span class="badge builder"></span> Dashboardv2RangeMapOptions
---
# <span class="badge builder"></span> Dashboardv2RangeMapOptions

## Constructor

```python
Dashboardv2RangeMapOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2.Dashboardv2RangeMapOptions
```

### <span class="badge object-method"></span> from_val

Min value of the range. It can be null which means -Infinity

```python
def from_val(from_val: float) -> typing.Self
```

### <span class="badge object-method"></span> result

Config to apply when the value is within the range

```python
def result(result: cogbuilder.Builder[dashboardv2.ValueMappingResult]) -> typing.Self
```

### <span class="badge object-method"></span> to

Max value of the range. It can be null which means +Infinity

```python
def to(to: float) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dashboardv2RangeMapOptions](./object-Dashboardv2RangeMapOptions.md)
