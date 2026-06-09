---
title: <span class="badge builder"></span> MetricsQueryOrLogsQueryOrAnnotationQueryBuilder
---
# <span class="badge builder"></span> MetricsQueryOrLogsQueryOrAnnotationQueryBuilder

## Constructor

```go
func NewMetricsQueryOrLogsQueryOrAnnotationQueryBuilder() *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AnnotationQuery

```go
func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) AnnotationQuery(annotationQuery cog.Builder[cloudwatch.AnnotationQuery]) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder
```

### <span class="badge object-method"></span> LogsQuery

```go
func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) LogsQuery(logsQuery cog.Builder[cloudwatch.LogsQuery]) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder
```

### <span class="badge object-method"></span> MetricsQuery

```go
func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) MetricsQuery(metricsQuery cog.Builder[cloudwatch.MetricsQuery]) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricsQueryOrLogsQueryOrAnnotationQuery](./object-MetricsQueryOrLogsQueryOrAnnotationQuery.md)
