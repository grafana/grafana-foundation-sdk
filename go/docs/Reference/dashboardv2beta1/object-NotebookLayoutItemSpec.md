---
title: <span class="badge object-type-struct"></span> NotebookLayoutItemSpec
---
# <span class="badge object-type-struct"></span> NotebookLayoutItemSpec

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Definition

```go
type NotebookLayoutItemSpec struct {
    Element dashboardv2beta1.ElementReference `json:"element"`
    Source dashboardv2beta1.NotebookLayoutItemSpecSource `json:"source"`
    Collapsed *bool `json:"collapsed,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotebookLayoutItemSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notebookLayoutItemSpec *NotebookLayoutItemSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotebookLayoutItemSpec` objects.

```go
func (notebookLayoutItemSpec *NotebookLayoutItemSpec) Equals(other NotebookLayoutItemSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotebookLayoutItemSpec` fields for violations and returns them.

```go
func (notebookLayoutItemSpec *NotebookLayoutItemSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotebookLayoutItemSpecBuilder](./builder-NotebookLayoutItemSpecBuilder.md)
