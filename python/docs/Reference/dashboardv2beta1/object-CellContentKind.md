---
title: <span class="badge object-type-disjunction"></span> CellContentKind
---
# <span class="badge object-type-disjunction"></span> CellContentKind

Pluggable cell content discriminated by `kind`. New content types are added

by extending this union with another <Name>CellContentKind member.

## Definition

```python
# Pluggable cell content discriminated by `kind`. New content types are added
# by extending this union with another <Name>CellContentKind member.
CellContentKind: typing.TypeAlias = typing.Union[dashboardv2beta1.MarkdownCellContentKind, dashboardv2beta1.CodeCellContentKind]
```
