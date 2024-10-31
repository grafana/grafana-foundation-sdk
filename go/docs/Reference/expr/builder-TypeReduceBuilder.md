---
title: <span class="badge builder"></span> TypeReduceBuilder
---
# <span class="badge builder"></span> TypeReduceBuilder

## Constructor

```go
func NewTypeReduceBuilder() *TypeReduceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TypeReduceBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Datasource

The datasource

```go
func (builder *TypeReduceBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeReduceBuilder
```

### <span class="badge object-method"></span> Expression

Reference to single query result

```go
func (builder *TypeReduceBuilder) Expression(expression string) *TypeReduceBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *TypeReduceBuilder) Hide(hide bool) *TypeReduceBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *TypeReduceBuilder) IntervalMs(intervalMs float64) *TypeReduceBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *TypeReduceBuilder) MaxDataPoints(maxDataPoints int64) *TypeReduceBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *TypeReduceBuilder) QueryType(queryType string) *TypeReduceBuilder
```

### <span class="badge object-method"></span> Reducer

The reducer

Possible enum values:

 - `"sum"` 

 - `"mean"` 

 - `"min"` 

 - `"max"` 

 - `"count"` 

 - `"last"` 

 - `"median"` 

```go
func (builder *TypeReduceBuilder) Reducer(reducer expr.TypeReduceReducer) *TypeReduceBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *TypeReduceBuilder) RefId(refId string) *TypeReduceBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *TypeReduceBuilder) ResultAssertions(resultAssertions cog.Builder[expr.ExprTypeReduceResultAssertions]) *TypeReduceBuilder
```

### <span class="badge object-method"></span> Settings

Reducer Options

```go
func (builder *TypeReduceBuilder) Settings(settings cog.Builder[expr.ExprTypeReduceSettings]) *TypeReduceBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *TypeReduceBuilder) TimeRange(timeRange cog.Builder[expr.ExprTypeReduceTimeRange]) *TypeReduceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TypeReduce](./object-TypeReduce.md)
