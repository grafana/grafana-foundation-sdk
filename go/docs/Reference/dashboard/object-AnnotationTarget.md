---
title: <span class="badge object-type-struct"></span> AnnotationTarget
---
# <span class="badge object-type-struct"></span> AnnotationTarget

TODO: this should be a regular DataQuery that depends on the selected dashboard

these match the properties of the "grafana" datasouce that is default in most dashboards

## Definition

```go
type AnnotationTarget struct {
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    Limit int64 `json:"limit"`
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    MatchAny bool `json:"matchAny"`
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    Tags []string `json:"tags"`
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationTarget` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationTarget *AnnotationTarget) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationTarget` objects.

```go
func (annotationTarget *AnnotationTarget) Equals(other AnnotationTarget) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationTarget` fields for violations and returns them.

```go
func (annotationTarget *AnnotationTarget) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationTargetBuilder](./builder-AnnotationTargetBuilder.md)
