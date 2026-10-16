---
title: <span class="badge object-type-class"></span> CellContentKind
---
# <span class="badge object-type-class"></span> CellContentKind

Pluggable cell content discriminated by `kind`. New content types are added

by extending this union with another <Name>CellContentKind member.

## Definition

```java
public class CellContentKind {
  public MarkdownCellContentKind markdownCellContentKind;
  public CodeCellContentKind codeCellContentKind;
}
```
