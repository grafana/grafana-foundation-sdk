---
title: <span class="badge builder"></span> RowsLayoutRow
---
# <span class="badge builder"></span> RowsLayoutRow

## Constructor

```python
RowsLayoutRow(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.RowsLayoutRowKind
```

### <span class="badge object-method"></span> auto_grid_layout

```python
def auto_grid_layout(auto_grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> rows_layout

```python
def rows_layout(rows_layout_kind: cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> tabs_layout

```python
def tabs_layout(tabs_layout_kind: cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]) -> typing.Self
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

### <span class="badge object-method"></span> grid_layout

```python
def grid_layout(grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.GridLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> hide_header

```python
def hide_header(hide_header: bool) -> typing.Self
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

 * <span class="badge object-type-class"></span> [RowsLayoutRowKind](./object-RowsLayoutRowKind.md)
