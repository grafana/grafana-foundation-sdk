---
title: <span class="badge object-type-interface"></span> NotebookLayoutItemSpec
---
# <span class="badge object-type-interface"></span> NotebookLayoutItemSpec

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Definition

```typescript
export interface NotebookLayoutItemSpec {
	element: dashboardv2beta1.ElementReference;
	source: "assistant" | "user";
	collapsed?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [NotebookLayoutItemSpecBuilder](./builder-NotebookLayoutItemSpecBuilder.md)
