---
title: <span class="badge object-type-struct"></span> CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery
---
# <span class="badge object-type-struct"></span> CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery

## Definition

```go
type CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery struct {
    CloudWatchMetricsQuery *cloudwatch.CloudWatchMetricsQuery `json:"CloudWatchMetricsQuery,omitempty"`
    CloudWatchLogsQuery *cloudwatch.CloudWatchLogsQuery `json:"CloudWatchLogsQuery,omitempty"`
    CloudWatchAnnotationQuery *cloudwatch.CloudWatchAnnotationQuery `json:"CloudWatchAnnotationQuery,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` as JSON.

```go
func (cloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` from JSON.

```go
func (cloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` objects.

```go
func (cloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) Equals(other CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` fields for violations and returns them.

```go
func (cloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) Validate() error
```

