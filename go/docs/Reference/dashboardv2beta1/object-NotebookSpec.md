---
title: <span class="badge object-type-struct"></span> NotebookSpec
---
# <span class="badge object-type-struct"></span> NotebookSpec

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Definition

```go
type NotebookSpec struct {
    Title string `json:"title"`
    Description *string `json:"description,omitempty"`
    Tags []string `json:"tags"`
    TimeSettings dashboardv2beta1.TimeSettingsSpec `json:"timeSettings"`
    Elements map[string]dashboardv2beta1.NotebookElement `json:"elements"`
    Layout dashboardv2beta1.NotebookLayoutKind `json:"layout"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotebookSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notebookSpec *NotebookSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotebookSpec` objects.

```go
func (notebookSpec *NotebookSpec) Equals(other NotebookSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotebookSpec` fields for violations and returns them.

```go
func (notebookSpec *NotebookSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotebookSpecBuilder](./builder-NotebookSpecBuilder.md)
