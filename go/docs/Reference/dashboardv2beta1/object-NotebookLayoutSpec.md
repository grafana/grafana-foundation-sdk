---
title: <span class="badge object-type-struct"></span> NotebookLayoutSpec
---
# <span class="badge object-type-struct"></span> NotebookLayoutSpec

## Definition

```go
type NotebookLayoutSpec struct {
    Cells []dashboardv2beta1.NotebookLayoutItemKind `json:"cells"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotebookLayoutSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notebookLayoutSpec *NotebookLayoutSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotebookLayoutSpec` objects.

```go
func (notebookLayoutSpec *NotebookLayoutSpec) Equals(other NotebookLayoutSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotebookLayoutSpec` fields for violations and returns them.

```go
func (notebookLayoutSpec *NotebookLayoutSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotebookLayoutSpecBuilder](./builder-NotebookLayoutSpecBuilder.md)
