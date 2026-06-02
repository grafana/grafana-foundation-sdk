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

### <span class="badge object-method"></span> BasicLogsQuery

If set to true the query will be run as a basic logs query

```go
func (builder *LogsQueryBuilder) BasicLogsQuery(basicLogsQuery bool) *LogsQueryBuilder
```

### <span class="badge object-method"></span> DashboardTime

If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.

```go
func (builder *LogsQueryBuilder) DashboardTime(dashboardTime bool) *LogsQueryBuilder
```

### <span class="badge object-method"></span> IntersectTime

@deprecated Use dashboardTime instead

```go
func (builder *LogsQueryBuilder) IntersectTime(intersectTime bool) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Query

KQL query to be executed.

```go
func (builder *LogsQueryBuilder) Query(query string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Resource

@deprecated Use resources instead

```go
func (builder *LogsQueryBuilder) Resource(resource string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Resources

Array of resource URIs to be queried.

```go
func (builder *LogsQueryBuilder) Resources(resources []string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> ResultFormat

Specifies the format results should be returned as.

```go
func (builder *LogsQueryBuilder) ResultFormat(resultFormat azuremonitor.ResultFormat) *LogsQueryBuilder
```

### <span class="badge object-method"></span> TimeColumn

If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated

```go
func (builder *LogsQueryBuilder) TimeColumn(timeColumn string) *LogsQueryBuilder
```

### <span class="badge object-method"></span> Workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat.

```go
func (builder *LogsQueryBuilder) Workspace(workspace string) *LogsQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LogsQuery](./object-LogsQuery.md)
