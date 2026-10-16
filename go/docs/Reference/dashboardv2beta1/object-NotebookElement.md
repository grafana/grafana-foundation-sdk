---
title: <span class="badge object-type-ref"></span> NotebookElement
---
# <span class="badge object-type-ref"></span> NotebookElement

A notebook element is a narrative cell, a panel, or a library panel. Unlike the dashboard

Element union, this one includes CellKind — and it is referenced ONLY by NotebookSpec.

CellKind is listed first so it is the generated default (a notebook is narrative-first).

## Definition

```go
type NotebookElement = dashboardv2beta1.CellKindOrV2PanelKindOrLibraryPanelKind
```
