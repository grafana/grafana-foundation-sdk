---
title: <span class="badge builder"></span> LineConfig
---
# <span class="badge builder"></span> LineConfig

## Constructor

```python
LineConfig()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> common.LineConfig
```

### <span class="badge object-method"></span> line_color

```python
def line_color(line_color: str) -> typing.Self
```

### <span class="badge object-method"></span> line_interpolation

```python
def line_interpolation(line_interpolation: common.LineInterpolation) -> typing.Self
```

### <span class="badge object-method"></span> line_style

```python
def line_style(line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self
```

### <span class="badge object-method"></span> line_width

```python
def line_width(line_width: float) -> typing.Self
```

### <span class="badge object-method"></span> span_nulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```python
def span_nulls(span_nulls: typing.Union[bool, float]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [LineConfig](./object-LineConfig.md)
