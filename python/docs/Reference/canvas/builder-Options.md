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
def build() -> canvas.Options
```

### <span class="badge object-method"></span> inline_editing

Enable inline editing

```python
def inline_editing(inline_editing: bool) -> typing.Self
```

### <span class="badge object-method"></span> root

The root element of canvas (frame), where all canvas elements are nested

TODO: Figure out how to define a default value for this

```python
def root(root: cogbuilder.Builder[canvas.CanvasOptionsRoot]) -> typing.Self
```

### <span class="badge object-method"></span> show_advanced_types

Show all available element types

```python
def show_advanced_types(show_advanced_types: bool) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
