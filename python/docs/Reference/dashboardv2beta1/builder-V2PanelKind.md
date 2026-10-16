---
title: <span class="badge builder"></span> V2PanelKind
---
# <span class="badge builder"></span> V2PanelKind

The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,

which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has

to be forked rather than shared because PanelKind reaches TransformationKind through

QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not

change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.

## Constructor

```python
V2PanelKind()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.V2PanelKind
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: cogbuilder.Builder[dashboardv2beta1.V2PanelSpec]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [V2PanelKind](./object-V2PanelKind.md)
