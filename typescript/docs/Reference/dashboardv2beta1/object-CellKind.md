---
title: <span class="badge object-type-interface"></span> CellKind
---
# <span class="badge object-type-interface"></span> CellKind

A cell holds non-panel narrative content (markdown text, code) in a notebook layout.

Panel cells are not represented here — they reuse V2PanelKind.

## Definition

```typescript
export interface CellKind {
	kind: "Cell";
	spec: dashboardv2beta1.CellSpec;
}

```
## See also

 * <span class="badge builder"></span> [CellKindBuilder](./builder-CellKindBuilder.md)
