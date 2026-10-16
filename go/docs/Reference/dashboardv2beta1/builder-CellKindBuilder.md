---
title: <span class="badge builder"></span> CellKindBuilder
---
# <span class="badge builder"></span> CellKindBuilder

A cell holds non-panel narrative content (markdown text, code) in a notebook layout.

Panel cells are not represented here — they reuse V2PanelKind.

## Constructor

```go
func NewCellKindBuilder() *CellKindBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CellKindBuilder) Build() (CellKind, error)
```

### <span class="badge object-method"></span> Spec

```go
func (builder *CellKindBuilder) Spec(spec cog.Builder[dashboardv2beta1.CellSpec]) *CellKindBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CellKind](./object-CellKind.md)
