---
title: <span class="badge object-type-struct"></span> MetricQuery
---
# <span class="badge object-type-struct"></span> MetricQuery

@deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.

## Definition

```go
type MetricQuery struct {
    // GCP project to execute the query against.
    ProjectName string `json:"projectName"`
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`
    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    AliasBy *string `json:"aliasBy,omitempty"`
    EditorMode string `json:"editorMode"`
    MetricType string `json:"metricType"`
    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    CrossSeriesReducer string `json:"crossSeriesReducer"`
    // Array of labels to group data by.
    GroupBys []string `json:"groupBys,omitempty"`
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    Filters []string `json:"filters,omitempty"`
    MetricKind *googlecloudmonitoring.MetricKind `json:"metricKind,omitempty"`
    ValueType *string `json:"valueType,omitempty"`
    View *string `json:"view,omitempty"`
    // MQL query to be executed.
    Query string `json:"query"`
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    Preprocessor *googlecloudmonitoring.PreprocessorType `json:"preprocessor,omitempty"`
    // To disable the graphPeriod, it should explictly be set to 'disabled'.
    GraphPeriod *string `json:"graphPeriod,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricQuery *MetricQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricQuery` objects.

```go
func (metricQuery *MetricQuery) Equals(other MetricQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricQuery` fields for violations and returns them.

```go
func (metricQuery *MetricQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
