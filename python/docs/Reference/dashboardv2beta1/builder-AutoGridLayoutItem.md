---
title: <span class="badge builder"></span> AutoGridLayoutItem
---
# <span class="badge builder"></span> AutoGridLayoutItem

## Constructor

```python
AutoGridLayoutItem(name: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.AutoGridLayoutItemKind
```

### <span class="badge object-method"></span> conditional_rendering

```python
def conditional_rendering(conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self
```

### <span class="badge object-method"></span> element

```python
def element(element: cogbuilder.Builder[dashboardv2beta1.ElementReference]) -> typing.Self
```

### <span class="badge object-method"></span> name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> repeat

```python
def repeat(repeat: cogbuilder.Builder[dashboardv2beta1.AutoGridRepeatOptions]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AutoGridLayoutItemKind](./object-AutoGridLayoutItemKind.md)
