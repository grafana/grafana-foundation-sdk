---
title: <span class="badge object-type-struct"></span> LegacyCloudMonitoringAnnotationQuery
---
# <span class="badge object-type-struct"></span> LegacyCloudMonitoringAnnotationQuery

@deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.

## Definition

```go
type LegacyCloudMonitoringAnnotationQuery struct {
    // GCP project to execute the query against.
    ProjectName string `json:"projectName"`
    MetricType string `json:"metricType"`
    // Query refId.
    RefId string `json:"refId"`
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    Filters []string `json:"filters"`
    MetricKind googlecloudmonitoring.MetricKind `json:"metricKind"`
    ValueType string `json:"valueType"`
    // Annotation title.
    Title string `json:"title"`
    // Annotation text.
    Text string `json:"text"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LegacyCloudMonitoringAnnotationQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (legacyCloudMonitoringAnnotationQuery *LegacyCloudMonitoringAnnotationQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LegacyCloudMonitoringAnnotationQuery` objects.

```go
func (legacyCloudMonitoringAnnotationQuery *LegacyCloudMonitoringAnnotationQuery) Equals(other LegacyCloudMonitoringAnnotationQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LegacyCloudMonitoringAnnotationQuery` fields for violations and returns them.

```go
func (legacyCloudMonitoringAnnotationQuery *LegacyCloudMonitoringAnnotationQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [LegacyCloudMonitoringAnnotationQueryBuilder](./builder-LegacyCloudMonitoringAnnotationQueryBuilder.md)
