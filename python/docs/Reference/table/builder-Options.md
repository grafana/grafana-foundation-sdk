---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> table.Options
```

### <span class="badge object-method"></span> cell_height

Controls the height of the rows

```python
def cell_height(cell_height: common.TableCellHeight) -> typing.Self
```

### <span class="badge object-method"></span> footer

Controls footer options

```python
def footer(footer: cogbuilder.Builder[common.TableFooterOptions]) -> typing.Self
```

### <span class="badge object-method"></span> frame_index

Represents the index of the selected frame

```python
def frame_index(frame_index: float) -> typing.Self
```

### <span class="badge object-method"></span> show_header

Controls whether the panel should show the header

```python
def show_header(show_header: bool) -> typing.Self
```

### <span class="badge object-method"></span> show_type_icons

Controls whether the header should show icons for the column types

```python
def show_type_icons(show_type_icons: bool) -> typing.Self
```

### <span class="badge object-method"></span> sort_by

Used to control row sorting

```python
def sort_by(sort_by: list[cogbuilder.Builder[common.TableSortByFieldState]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
