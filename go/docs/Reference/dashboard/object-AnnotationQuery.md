---
title: <span class="badge object-type-struct"></span> AnnotationQuery
---
# <span class="badge object-type-struct"></span> AnnotationQuery

TODO docs

FROM: AnnotationQuery in grafana-data/src/types/annotations.ts

## Definition

```go
type AnnotationQuery struct {
    // Name of annotation.
    Name string `json:"name"`
    // Datasource where the annotations data is
    Datasource dashboard.DataSourceRef `json:"datasource"`
    // When enabled the annotation query is issued with every dashboard refresh
    Enable bool `json:"enable"`
    // Annotation queries can be toggled on or off at the top of the dashboard.
    // When hide is true, the toggle is not shown in the dashboard.
    Hide *bool `json:"hide,omitempty"`
    // Color to use for the annotation event markers
    IconColor string `json:"iconColor"`
    // Filters to apply when fetching annotations
    Filter *dashboard.AnnotationPanelFilter `json:"filter,omitempty"`
    // TODO.. this should just be a normal query target
    Target *dashboard.AnnotationTarget `json:"target,omitempty"`
    // TODO -- this should not exist here, it is based on the --grafana-- datasource
    Type *string `json:"type,omitempty"`
    Expr *string `json:"expr,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationQuery *AnnotationQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationQuery` objects.

```go
func (annotationQuery *AnnotationQuery) Equals(other AnnotationQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationQuery` fields for violations and returns them.

```go
func (annotationQuery *AnnotationQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
