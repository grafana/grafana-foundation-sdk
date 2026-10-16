---
title: <span class="badge builder"></span> V2TransformationKind
---
# <span class="badge builder"></span> V2TransformationKind

Dashboard v2 shape: the transformation ID moved from `kind` to `group`.

## Constructor

```python
V2TransformationKind()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.V2TransformationKind
```

### <span class="badge object-method"></span> group

The group is the transformation ID

```python
def group(group: str) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: cogbuilder.Builder[dashboardv2beta1.V2TransformationSpec]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [V2TransformationKind](./object-V2TransformationKind.md)
