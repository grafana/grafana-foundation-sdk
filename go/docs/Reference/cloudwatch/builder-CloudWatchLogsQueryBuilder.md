---
title: <span class="badge builder"></span> CloudWatchLogsQueryBuilder
---
# <span class="badge builder"></span> CloudWatchLogsQueryBuilder

## Constructor

```go
func NewCloudWatchLogsQueryBuilder() *CloudWatchLogsQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CloudWatchLogsQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *CloudWatchLogsQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> Expression

The CloudWatch Logs Insights query to execute

```go
func (builder *CloudWatchLogsQueryBuilder) Expression(expression string) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *CloudWatchLogsQueryBuilder) Hide(hide bool) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *CloudWatchLogsQueryBuilder) Id(id string) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> LogGroupNames

@deprecated use logGroups

```go
func (builder *CloudWatchLogsQueryBuilder) LogGroupNames(logGroupNames []string) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> LogGroups

Log groups to query

```go
func (builder *CloudWatchLogsQueryBuilder) LogGroups(logGroups []cog.Builder[cloudwatch.LogGroup]) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> QueryMode

Whether a query is a Metrics, Logs, or Annotations query

```go
func (builder *CloudWatchLogsQueryBuilder) QueryMode(queryMode cloudwatch.CloudWatchQueryMode) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *CloudWatchLogsQueryBuilder) QueryType(queryType string) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *CloudWatchLogsQueryBuilder) RefId(refId string) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> Region

AWS region to query for the logs

```go
func (builder *CloudWatchLogsQueryBuilder) Region(region string) *CloudWatchLogsQueryBuilder
```

### <span class="badge object-method"></span> StatsGroups

Fields to group the results by, this field is automatically populated whenever the query is updated

```go
func (builder *CloudWatchLogsQueryBuilder) StatsGroups(statsGroups []string) *CloudWatchLogsQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CloudWatchLogsQuery](./object-CloudWatchLogsQuery.md)
