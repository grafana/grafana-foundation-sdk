---
title: <span class="badge builder"></span> CanvasElementOptions
---
# <span class="badge builder"></span> CanvasElementOptions

## Constructor

```python
CanvasElementOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> canvas.CanvasElementOptions
```

### <span class="badge object-method"></span> background

```python
def background(background: cogbuilder.Builder[canvas.BackgroundConfig]) -> typing.Self
```

### <span class="badge object-method"></span> border

```python
def border(border: cogbuilder.Builder[canvas.LineConfig]) -> typing.Self
```

### <span class="badge object-method"></span> config

TODO: figure out how to define this (element config(s))

```python
def config(config: object) -> typing.Self
```

### <span class="badge object-method"></span> connections

```python
def connections(connections: list[cogbuilder.Builder[canvas.CanvasConnection]]) -> typing.Self
```

### <span class="badge object-method"></span> constraint

```python
def constraint(constraint: cogbuilder.Builder[canvas.Constraint]) -> typing.Self
```

### <span class="badge object-method"></span> name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> placement

```python
def placement(placement: cogbuilder.Builder[canvas.Placement]) -> typing.Self
```

### <span class="badge object-method"></span> type_val

```python
def type_val(type_val: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CanvasElementOptions](./object-CanvasElementOptions.md)
