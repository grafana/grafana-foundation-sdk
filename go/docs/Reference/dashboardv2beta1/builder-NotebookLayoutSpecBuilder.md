---
title: <span class="badge builder"></span> NotebookLayoutSpecBuilder
---
# <span class="badge builder"></span> NotebookLayoutSpecBuilder

## Constructor

```go
func NewNotebookLayoutSpecBuilder() *NotebookLayoutSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NotebookLayoutSpecBuilder) Build() (NotebookLayoutSpec, error)
```

### <span class="badge object-method"></span> Cells

```go
func (builder *NotebookLayoutSpecBuilder) Cells(cells []cog.Builder[dashboardv2beta1.NotebookLayoutItemKind]) *NotebookLayoutSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NotebookLayoutSpec](./object-NotebookLayoutSpec.md)
