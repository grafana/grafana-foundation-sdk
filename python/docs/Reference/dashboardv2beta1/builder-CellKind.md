---
title: <span class="badge builder"></span> CellKind
---
# <span class="badge builder"></span> CellKind

A cell holds non-panel narrative content (markdown text, code) in a notebook layout.

Panel cells are not represented here — they reuse V2PanelKind.

## Constructor

```python
CellKind()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.CellKind
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: cogbuilder.Builder[dashboardv2beta1.CellSpec]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CellKind](./object-CellKind.md)
