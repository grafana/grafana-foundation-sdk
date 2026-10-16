---
title: <span class="badge object-type-struct"></span> NotebookLayoutItemKind
---
# <span class="badge object-type-struct"></span> NotebookLayoutItemKind

## Definition

```go
type NotebookLayoutItemKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.NotebookLayoutItemSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotebookLayoutItemKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notebookLayoutItemKind *NotebookLayoutItemKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotebookLayoutItemKind` objects.

```go
func (notebookLayoutItemKind *NotebookLayoutItemKind) Equals(other NotebookLayoutItemKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotebookLayoutItemKind` fields for violations and returns them.

```go
func (notebookLayoutItemKind *NotebookLayoutItemKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotebookLayoutItemKindBuilder](./builder-NotebookLayoutItemKindBuilder.md)
