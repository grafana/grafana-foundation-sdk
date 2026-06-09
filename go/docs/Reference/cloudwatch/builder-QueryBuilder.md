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

### <span class="badge object-method"></span> AnnotationQuery

```go
func (builder *QueryBuilder) AnnotationQuery(annotationQuery cog.Builder[cloudwatch.AnnotationQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> LogsQuery

```go
func (builder *QueryBuilder) LogsQuery(logsQuery cog.Builder[cloudwatch.LogsQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> MetricsQuery

```go
func (builder *QueryBuilder) MetricsQuery(metricsQuery cog.Builder[cloudwatch.MetricsQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryBuilder) Labels(labels map[string]string) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
