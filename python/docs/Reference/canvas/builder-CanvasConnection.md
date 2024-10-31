---
title: <span class="badge builder"></span> CanvasConnection
---
# <span class="badge builder"></span> CanvasConnection

## Constructor

```python
CanvasConnection()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> canvas.CanvasConnection
```

### <span class="badge object-method"></span> color

```python
def color(color: cogbuilder.Builder[common.ColorDimensionConfig]) -> typing.Self
```

### <span class="badge object-method"></span> path

```python
def path(path: canvas.ConnectionPath) -> typing.Self
```

### <span class="badge object-method"></span> size

```python
def size(size: cogbuilder.Builder[common.ScaleDimensionConfig]) -> typing.Self
```

### <span class="badge object-method"></span> source

```python
def source(source: cogbuilder.Builder[canvas.ConnectionCoordinates]) -> typing.Self
```

### <span class="badge object-method"></span> source_original

```python
def source_original(source_original: cogbuilder.Builder[canvas.ConnectionCoordinates]) -> typing.Self
```

### <span class="badge object-method"></span> target

```python
def target(target: cogbuilder.Builder[canvas.ConnectionCoordinates]) -> typing.Self
```

### <span class="badge object-method"></span> target_name

```python
def target_name(target_name: str) -> typing.Self
```

### <span class="badge object-method"></span> target_original

```python
def target_original(target_original: cogbuilder.Builder[canvas.ConnectionCoordinates]) -> typing.Self
```

### <span class="badge object-method"></span> vertices

```python
def vertices(vertices: list[cogbuilder.Builder[canvas.ConnectionCoordinates]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CanvasConnection](./object-CanvasConnection.md)
