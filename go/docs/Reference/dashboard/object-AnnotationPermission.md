---
title: <span class="badge object-type-struct"></span> AnnotationPermission
---
# <span class="badge object-type-struct"></span> AnnotationPermission

## Definition

```go
type AnnotationPermission struct {
    Dashboard *dashboard.AnnotationActions `json:"dashboard,omitempty"`
    Organization *dashboard.AnnotationActions `json:"organization,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationPermission` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationPermission *AnnotationPermission) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationPermission` objects.

```go
func (annotationPermission *AnnotationPermission) Equals(other AnnotationPermission) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationPermission` fields for violations and returns them.

```go
func (annotationPermission *AnnotationPermission) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationPermissionBuilder](./builder-AnnotationPermissionBuilder.md)
