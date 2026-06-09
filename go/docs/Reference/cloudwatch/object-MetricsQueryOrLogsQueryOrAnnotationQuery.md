---
title: <span class="badge object-type-struct"></span> MetricsQueryOrLogsQueryOrAnnotationQuery
---
# <span class="badge object-type-struct"></span> MetricsQueryOrLogsQueryOrAnnotationQuery

## Definition

```go
type MetricsQueryOrLogsQueryOrAnnotationQuery struct {
    MetricsQuery *cloudwatch.MetricsQuery `json:"MetricsQuery,omitempty"`
    LogsQuery *cloudwatch.LogsQuery `json:"LogsQuery,omitempty"`
    AnnotationQuery *cloudwatch.AnnotationQuery `json:"AnnotationQuery,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `MetricsQueryOrLogsQueryOrAnnotationQuery` as JSON.

```go
func (metricsQueryOrLogsQueryOrAnnotationQuery *MetricsQueryOrLogsQueryOrAnnotationQuery) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `MetricsQueryOrLogsQueryOrAnnotationQuery` from JSON.

```go
func (metricsQueryOrLogsQueryOrAnnotationQuery *MetricsQueryOrLogsQueryOrAnnotationQuery) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricsQueryOrLogsQueryOrAnnotationQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricsQueryOrLogsQueryOrAnnotationQuery *MetricsQueryOrLogsQueryOrAnnotationQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricsQueryOrLogsQueryOrAnnotationQuery` objects.

```go
func (metricsQueryOrLogsQueryOrAnnotationQuery *MetricsQueryOrLogsQueryOrAnnotationQuery) Equals(other MetricsQueryOrLogsQueryOrAnnotationQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricsQueryOrLogsQueryOrAnnotationQuery` fields for violations and returns them.

```go
func (metricsQueryOrLogsQueryOrAnnotationQuery *MetricsQueryOrLogsQueryOrAnnotationQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricsQueryOrLogsQueryOrAnnotationQueryBuilder](./builder-MetricsQueryOrLogsQueryOrAnnotationQueryBuilder.md)
