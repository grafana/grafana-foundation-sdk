---
title: <span class="badge object-type-struct"></span> AnnotationContainer
---
# <span class="badge object-type-struct"></span> AnnotationContainer

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

## Definition

```go
type AnnotationContainer struct {
    // List of annotations
    List []dashboard.AnnotationQuery `json:"list,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationContainer` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationContainer *AnnotationContainer) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationContainer` objects.

```go
func (annotationContainer *AnnotationContainer) Equals(other AnnotationContainer) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationContainer` fields for violations and returns them.

```go
func (annotationContainer *AnnotationContainer) Validate() error
```

