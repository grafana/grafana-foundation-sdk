---
title: <span class="badge object-type-struct"></span> AnnotationQueryKind
---
# <span class="badge object-type-struct"></span> AnnotationQueryKind

## Definition

```go
type AnnotationQueryKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.AnnotationQuerySpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationQueryKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationQueryKind *AnnotationQueryKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationQueryKind` objects.

```go
func (annotationQueryKind *AnnotationQueryKind) Equals(other AnnotationQueryKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationQueryKind` fields for violations and returns them.

```go
func (annotationQueryKind *AnnotationQueryKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
