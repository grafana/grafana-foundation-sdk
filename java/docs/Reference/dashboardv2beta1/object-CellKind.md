---
title: <span class="badge object-type-class"></span> CellKind
---
# <span class="badge object-type-class"></span> CellKind

A cell holds non-panel narrative content (markdown text, code) in a notebook layout.

Panel cells are not represented here — they reuse V2PanelKind.

## Definition

```java
public class CellKind {
  public String kind;
  public CellSpec spec;
}
```
## See also

 * <span class="badge builder"></span> [CellKindBuilder](./builder-CellKindBuilder.md)
