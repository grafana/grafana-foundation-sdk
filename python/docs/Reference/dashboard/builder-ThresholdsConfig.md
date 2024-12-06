---
title: <span class="badge builder"></span> ThresholdsConfig
---
# <span class="badge builder"></span> ThresholdsConfig

## Constructor

```python
ThresholdsConfig()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.ThresholdsConfig
```

### <span class="badge object-method"></span> mode

Thresholds mode.

```python
def mode(mode: dashboard.ThresholdsMode) -> typing.Self
```

### <span class="badge object-method"></span> steps

Must be sorted by 'value', first value is always -Infinity

```python
def steps(steps: list[dashboard.Threshold]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [ThresholdsConfig](./object-ThresholdsConfig.md)
