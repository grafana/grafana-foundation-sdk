---
title: <span class="badge builder"></span> AutoGridLayoutSpec
---
# <span class="badge builder"></span> AutoGridLayoutSpec

## Constructor

```python
AutoGridLayoutSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.AutoGridLayoutSpec
```

### <span class="badge object-method"></span> column_width

```python
def column_width(column_width: float) -> typing.Self
```

### <span class="badge object-method"></span> column_width_mode

```python
def column_width_mode(column_width_mode: typing.Literal["narrow", "standard", "wide", "custom"]) -> typing.Self
```

### <span class="badge object-method"></span> fill_screen

```python
def fill_screen(fill_screen: bool) -> typing.Self
```

### <span class="badge object-method"></span> items

```python
def items(items: list[cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutItemKind]]) -> typing.Self
```

### <span class="badge object-method"></span> max_column_count

```python
def max_column_count(max_column_count: float) -> typing.Self
```

### <span class="badge object-method"></span> row_height

```python
def row_height(row_height: float) -> typing.Self
```

### <span class="badge object-method"></span> row_height_mode

```python
def row_height_mode(row_height_mode: typing.Literal["short", "standard", "tall", "custom"]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AutoGridLayoutSpec](./object-AutoGridLayoutSpec.md)
