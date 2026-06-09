---
title: <span class="badge builder"></span> LogsQueryBuilder
---
# <span class="badge builder"></span> LogsQueryBuilder

## Constructor

```go
func NewLogsQueryBuilder() *LogsQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LogsQueryBuilder) Build() (LogsQuery, error)
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *LogsQueryBuilder) Datasource(datasource common.DataSourceRef) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Expression

The CloudWatch Logs Insights query to execute

```go
func (builder *LogsQueryBuilder) Expression(expression string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *LogsQueryBuilder) Hide(hide bool) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *LogsQueryBuilder) Id(id string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> LogGroupNames

@deprecated use logGroups

```go
func (builder *LogsQueryBuilder) LogGroupNames(logGroupNames []string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> LogGroups

Log groups to query

```go
func (builder *LogsQueryBuilder) LogGroups(logGroups []cog.Builder[cloudwatch.LogGroup]) *LogsQueryBuilder
```

### <span class="badge object-method"></span> QueryLanguage

Language used for querying logs, can be CWLI, SQL, or PPL. If empty, the default language is CWLI.

```go
func (builder *LogsQueryBuilder) QueryLanguage(queryLanguage cloudwatch.LogsQueryLanguage) *LogsQueryBuilder
```

### <span class="badge object-method"></span> QueryMode

Whether a query is a Metrics, Logs, or Annotations query

```go
func (builder *LogsQueryBuilder) QueryMode(queryMode cloudwatch.QueryMode) *LogsQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *LogsQueryBuilder) QueryType(queryType string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *LogsQueryBuilder) RefId(refId string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Region

AWS region to query for the logs

```go
func (builder *LogsQueryBuilder) Region(region string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> StatsGroups

Fields to group the results by, this field is automatically populated whenever the query is updated

```go
func (builder *LogsQueryBuilder) StatsGroups(statsGroups []string) *LogsQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LogsQuery](./object-LogsQuery.md)
