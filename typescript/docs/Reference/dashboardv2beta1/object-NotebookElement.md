---
title: <span class="badge object-type-disjunction"></span> NotebookElement
---
# <span class="badge object-type-disjunction"></span> NotebookElement

A notebook element is a narrative cell, a panel, or a library panel. Unlike the dashboard

Element union, this one includes CellKind — and it is referenced ONLY by NotebookSpec.

CellKind is listed first so it is the generated default (a notebook is narrative-first).

## Definition

```typescript
export type NotebookElement = dashboardv2beta1.CellKind | dashboardv2beta1.V2PanelKind | dashboardv2beta1.LibraryPanelKind;

```
