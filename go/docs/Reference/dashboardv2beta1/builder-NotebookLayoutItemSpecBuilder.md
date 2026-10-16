---
title: <span class="badge builder"></span> NotebookLayoutItemSpecBuilder
---
# <span class="badge builder"></span> NotebookLayoutItemSpecBuilder

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Constructor

```go
func NewNotebookLayoutItemSpecBuilder() *NotebookLayoutItemSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NotebookLayoutItemSpecBuilder) Build() (NotebookLayoutItemSpec, error)
```

### <span class="badge object-method"></span> Collapsed

```go
func (builder *NotebookLayoutItemSpecBuilder) Collapsed(collapsed bool) *NotebookLayoutItemSpecBuilder
```

### <span class="badge object-method"></span> Element

```go
func (builder *NotebookLayoutItemSpecBuilder) Element(element cog.Builder[dashboardv2beta1.ElementReference]) *NotebookLayoutItemSpecBuilder
```

### <span class="badge object-method"></span> Source

```go
func (builder *NotebookLayoutItemSpecBuilder) Source(source dashboardv2beta1.NotebookLayoutItemSpecSource) *NotebookLayoutItemSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NotebookLayoutItemSpec](./object-NotebookLayoutItemSpec.md)
