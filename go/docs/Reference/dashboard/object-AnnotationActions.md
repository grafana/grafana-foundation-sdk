---
title: <span class="badge object-type-struct"></span> AnnotationActions
---
# <span class="badge object-type-struct"></span> AnnotationActions

## Definition

```go
type AnnotationActions struct {
    CanAdd *bool `json:"canAdd,omitempty"`
    CanDelete *bool `json:"canDelete,omitempty"`
    CanEdit *bool `json:"canEdit,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationActions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationActions *AnnotationActions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationActions` objects.

```go
func (annotationActions *AnnotationActions) Equals(other AnnotationActions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationActions` fields for violations and returns them.

```go
func (annotationActions *AnnotationActions) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationActionsBuilder](./builder-AnnotationActionsBuilder.md)
