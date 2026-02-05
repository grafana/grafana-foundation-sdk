---
title: <span class="badge builder"></span> RowsLayoutRowSpec
---
# <span class="badge builder"></span> RowsLayoutRowSpec

## Constructor

```python
RowsLayoutRowSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.RowsLayoutRowSpec
```

### <span class="badge object-method"></span> collapse

```python
def collapse(collapse: bool) -> typing.Self
```

### <span class="badge object-method"></span> conditional_rendering

```python
def conditional_rendering(conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self
```

### <span class="badge object-method"></span> fill_screen

```python
def fill_screen(fill_screen: bool) -> typing.Self
```

### <span class="badge object-method"></span> hide_header

```python
def hide_header(hide_header: bool) -> typing.Self
```

### <span class="badge object-method"></span> layout

```python
def layout(layout: typing.Union[cogbuilder.Builder[dashboardv2beta1.GridLayoutKind], cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind], cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind], cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]]) -> typing.Self
```

### <span class="badge object-method"></span> repeat

```python
def repeat(repeat: cogbuilder.Builder[dashboardv2beta1.RowRepeatOptions]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RowsLayoutRowSpec](./object-RowsLayoutRowSpec.md)
