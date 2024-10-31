---
title: <span class="badge builder"></span> TypeSqlBuilder
---
# <span class="badge builder"></span> TypeSqlBuilder

## Constructor

```go
func NewTypeSqlBuilder() *TypeSqlBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TypeSqlBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Datasource

The datasource

```go
func (builder *TypeSqlBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeSqlBuilder
```

### <span class="badge object-method"></span> Expression

```go
func (builder *TypeSqlBuilder) Expression(expression string) *TypeSqlBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *TypeSqlBuilder) Hide(hide bool) *TypeSqlBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *TypeSqlBuilder) IntervalMs(intervalMs float64) *TypeSqlBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *TypeSqlBuilder) MaxDataPoints(maxDataPoints int64) *TypeSqlBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *TypeSqlBuilder) QueryType(queryType string) *TypeSqlBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *TypeSqlBuilder) RefId(refId string) *TypeSqlBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *TypeSqlBuilder) ResultAssertions(resultAssertions cog.Builder[expr.ExprTypeSqlResultAssertions]) *TypeSqlBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *TypeSqlBuilder) TimeRange(timeRange cog.Builder[expr.ExprTypeSqlTimeRange]) *TypeSqlBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TypeSql](./object-TypeSql.md)
