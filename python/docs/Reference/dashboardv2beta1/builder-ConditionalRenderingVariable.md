---
title: <span class="badge builder"></span> ConditionalRenderingVariable
---
# <span class="badge builder"></span> ConditionalRenderingVariable

## Constructor

```python
ConditionalRenderingVariable()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.ConditionalRenderingVariableKind
```

### <span class="badge object-method"></span> operator

```python
def operator(operator: typing.Literal["equals", "notEquals", "matches", "notMatches"]) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableSpec]) -> typing.Self
```

### <span class="badge object-method"></span> value

```python
def value(value: str) -> typing.Self
```

### <span class="badge object-method"></span> variable

```python
def variable(variable: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [ConditionalRenderingVariableKind](./object-ConditionalRenderingVariableKind.md)
