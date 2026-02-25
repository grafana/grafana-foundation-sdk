---
title: <span class="badge object-type-struct"></span> AnnotationQuerySpec
---
# <span class="badge object-type-struct"></span> AnnotationQuerySpec

## Definition

```go
type AnnotationQuerySpec struct {
    Query dashboardv2beta1.DataQueryKind `json:"query"`
    Enable bool `json:"enable"`
    Hide bool `json:"hide"`
    IconColor string `json:"iconColor"`
    Name string `json:"name"`
    BuiltIn *bool `json:"builtIn,omitempty"`
    Filter *dashboardv2beta1.AnnotationPanelFilter `json:"filter,omitempty"`
    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    Placement *string `json:"placement,omitempty"`
    // Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    LegacyOptions map[string]any `json:"legacyOptions,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationQuerySpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationQuerySpec *AnnotationQuerySpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationQuerySpec` objects.

```go
func (annotationQuerySpec *AnnotationQuerySpec) Equals(other AnnotationQuerySpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationQuerySpec` fields for violations and returns them.

```go
func (annotationQuerySpec *AnnotationQuerySpec) Validate() error
```

