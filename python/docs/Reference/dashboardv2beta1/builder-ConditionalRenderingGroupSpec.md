---
title: <span class="badge builder"></span> ConditionalRenderingGroupSpec
---
# <span class="badge builder"></span> ConditionalRenderingGroupSpec

## Constructor

```python
ConditionalRenderingGroupSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.ConditionalRenderingGroupSpec
```

### <span class="badge object-method"></span> condition

```python
def condition(condition: typing.Literal["and", "or"]) -> typing.Self
```

### <span class="badge object-method"></span> items

```python
def items(items: list[typing.Union[cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableKind], cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingDataKind], cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind]]]) -> typing.Self
```

### <span class="badge object-method"></span> visibility

```python
def visibility(visibility: typing.Literal["show", "hide"]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [ConditionalRenderingGroupSpec](./object-ConditionalRenderingGroupSpec.md)
