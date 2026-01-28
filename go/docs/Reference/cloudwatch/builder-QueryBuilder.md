---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```go
func NewQueryBuilder() *QueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error)
```

### <span class="badge object-method"></span> CloudWatchAnnotationQuery

```go
func (builder *QueryBuilder) CloudWatchAnnotationQuery(cloudWatchAnnotationQuery cog.Builder[cloudwatch.CloudWatchAnnotationQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> CloudWatchLogsQuery

```go
func (builder *QueryBuilder) CloudWatchLogsQuery(cloudWatchLogsQuery cog.Builder[cloudwatch.CloudWatchLogsQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> CloudWatchMetricsQuery

```go
func (builder *QueryBuilder) CloudWatchMetricsQuery(cloudWatchMetricsQuery cog.Builder[cloudwatch.CloudWatchMetricsQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
