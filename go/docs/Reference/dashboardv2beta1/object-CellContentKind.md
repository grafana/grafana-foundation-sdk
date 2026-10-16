---
title: <span class="badge object-type-ref"></span> CellContentKind
---
# <span class="badge object-type-ref"></span> CellContentKind

Pluggable cell content discriminated by `kind`. New content types are added

by extending this union with another <Name>CellContentKind member.

## Definition

```go
type CellContentKind = dashboardv2beta1.MarkdownCellContentKindOrCodeCellContentKind
```
