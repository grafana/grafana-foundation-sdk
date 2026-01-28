---
title: <span class="badge builder"></span> TabsLayoutTabSpec
---
# <span class="badge builder"></span> TabsLayoutTabSpec

## Constructor

```python
TabsLayoutTabSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.TabsLayoutTabSpec
```

### <span class="badge object-method"></span> conditional_rendering

```python
def conditional_rendering(conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self
```

### <span class="badge object-method"></span> layout

```python
def layout(layout: typing.Union[cogbuilder.Builder[dashboardv2beta1.GridLayoutKind], cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind], cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind], cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]]) -> typing.Self
```

### <span class="badge object-method"></span> repeat

```python
def repeat(repeat: cogbuilder.Builder[dashboardv2beta1.TabRepeatOptions]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TabsLayoutTabSpec](./object-TabsLayoutTabSpec.md)
