---
title: <span class="badge object-type-struct"></span> AnnotationQuery
---
# <span class="badge object-type-struct"></span> AnnotationQuery

Annotation sub-query properties.

## Definition

```go
type AnnotationQuery struct {
    // GCP project to execute the query against.
    ProjectName string `json:"projectName"`
    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    CrossSeriesReducer string `json:"crossSeriesReducer"`
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
    // Array of labels to group data by.
    GroupBys []string `json:"groupBys,omitempty"`
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    Filters []string `json:"filters,omitempty"`
    // Data view, defaults to FULL.
    View *string `json:"view,omitempty"`
    // Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    SecondaryCrossSeriesReducer *string `json:"secondaryCrossSeriesReducer,omitempty"`
    // Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    SecondaryAlignmentPeriod *string `json:"secondaryAlignmentPeriod,omitempty"`
    // Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
    SecondaryPerSeriesAligner *string `json:"secondaryPerSeriesAligner,omitempty"`
    // Only present if a preprocessor is selected. Array of labels to group data by.
    SecondaryGroupBys []string `json:"secondaryGroupBys,omitempty"`
    // Annotation title.
    Title *string `json:"title,omitempty"`
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    Preprocessor *googlecloudmonitoring.PreprocessorType `json:"preprocessor,omitempty"`
    // Annotation text.
    Text *string `json:"text,omitempty"`
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
