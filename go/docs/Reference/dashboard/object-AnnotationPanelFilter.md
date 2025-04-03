---
title: <span class="badge object-type-struct"></span> AnnotationPanelFilter
---
# <span class="badge object-type-struct"></span> AnnotationPanelFilter

## Definition

```go
type AnnotationPanelFilter struct {
    // Should the specified panels be included or excluded
    Exclude *bool `json:"exclude,omitempty"`
    // Panel IDs that should be included or excluded
    Ids []uint8 `json:"ids"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationPanelFilter` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationPanelFilter *AnnotationPanelFilter) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationPanelFilter` objects.

```go
func (annotationPanelFilter *AnnotationPanelFilter) Equals(other AnnotationPanelFilter) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationPanelFilter` fields for violations and returns them.

```go
func (annotationPanelFilter *AnnotationPanelFilter) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationPanelFilterBuilder](./builder-AnnotationPanelFilterBuilder.md)
