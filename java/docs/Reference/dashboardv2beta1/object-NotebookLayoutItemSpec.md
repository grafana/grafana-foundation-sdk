---
title: <span class="badge object-type-class"></span> NotebookLayoutItemSpec
---
# <span class="badge object-type-class"></span> NotebookLayoutItemSpec

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Definition

```java
public class NotebookLayoutItemSpec {
  public ElementReference element;
  public NotebookLayoutItemSpecSource source;
  public Boolean collapsed;
}
```
## See also

 * <span class="badge builder"></span> [NotebookLayoutItemSpecBuilder](./builder-NotebookLayoutItemSpecBuilder.md)
