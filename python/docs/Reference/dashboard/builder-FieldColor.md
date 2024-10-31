---
title: <span class="badge builder"></span> FieldColor
---
# <span class="badge builder"></span> FieldColor

## Constructor

```python
FieldColor()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.FieldColor
```

### <span class="badge object-method"></span> fixed_color

The fixed color value for fixed or shades color modes.

```python
def fixed_color(fixed_color: str) -> typing.Self
```

### <span class="badge object-method"></span> mode

The main color scheme mode.

```python
def mode(mode: dashboard.FieldColorModeId) -> typing.Self
```

### <span class="badge object-method"></span> series_by

Some visualizations need to know how to assign a series color from by value color schemes.

```python
def series_by(series_by: dashboard.FieldColorSeriesByMode) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [FieldColor](./object-FieldColor.md)
