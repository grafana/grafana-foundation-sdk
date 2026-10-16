---
title: <span class="badge object-type-class"></span> NotebookElement
---
# <span class="badge object-type-class"></span> NotebookElement

A notebook element is a narrative cell, a panel, or a library panel. Unlike the dashboard

Element union, this one includes CellKind — and it is referenced ONLY by NotebookSpec.

CellKind is listed first so it is the generated default (a notebook is narrative-first).

## Definition

```java
public class NotebookElement {
  public CellKind cellKind;
  public V2PanelKind v2PanelKind;
  public LibraryPanelKind libraryPanelKind;
}
```
