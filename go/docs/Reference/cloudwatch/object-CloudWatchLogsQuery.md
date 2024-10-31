---
title: <span class="badge object-type-struct"></span> CloudWatchLogsQuery
---
# <span class="badge object-type-struct"></span> CloudWatchLogsQuery

Shape of a CloudWatch Logs query

## Definition

```go
type CloudWatchLogsQuery struct {
    // Whether a query is a Metrics, Logs, or Annotations query
    QueryMode cloudwatch.CloudWatchQueryMode `json:"queryMode"`
    Id string `json:"id"`
    // AWS region to query for the logs
    Region string `json:"region"`
    // The CloudWatch Logs Insights query to execute
    Expression *string `json:"expression,omitempty"`
    // Fields to group the results by, this field is automatically populated whenever the query is updated
    StatsGroups []string `json:"statsGroups,omitempty"`
    // Log groups to query
    LogGroups []cloudwatch.LogGroup `json:"logGroups,omitempty"`
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    RefId string `json:"refId"`
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    Hide *bool `json:"hide,omitempty"`
    // Specify the query flavor
    // TODO make this required and give it a default
    QueryType *string `json:"queryType,omitempty"`
    // @deprecated use logGroups
    LogGroupNames []string `json:"logGroupNames,omitempty"`
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchLogsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cloudWatchLogsQuery *CloudWatchLogsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CloudWatchLogsQuery` objects.

```go
func (cloudWatchLogsQuery *CloudWatchLogsQuery) Equals(other CloudWatchLogsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CloudWatchLogsQuery` fields for violations and returns them.

```go
func (cloudWatchLogsQuery *CloudWatchLogsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [CloudWatchLogsQueryBuilder](./builder-CloudWatchLogsQueryBuilder.md)
