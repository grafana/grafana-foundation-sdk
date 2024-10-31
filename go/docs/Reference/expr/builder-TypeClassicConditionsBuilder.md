---
title: <span class="badge builder"></span> TypeClassicConditionsBuilder
---
# <span class="badge builder"></span> TypeClassicConditionsBuilder

## Constructor

```go
func NewTypeClassicConditionsBuilder() *TypeClassicConditionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TypeClassicConditionsBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Conditions

```go
func (builder *TypeClassicConditionsBuilder) Conditions(conditions []cog.Builder[expr.ExprTypeClassicConditionsConditions]) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> Datasource

The datasource

```go
func (builder *TypeClassicConditionsBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *TypeClassicConditionsBuilder) Hide(hide bool) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *TypeClassicConditionsBuilder) IntervalMs(intervalMs float64) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *TypeClassicConditionsBuilder) MaxDataPoints(maxDataPoints int64) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *TypeClassicConditionsBuilder) QueryType(queryType string) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *TypeClassicConditionsBuilder) RefId(refId string) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *TypeClassicConditionsBuilder) ResultAssertions(resultAssertions cog.Builder[expr.ExprTypeClassicConditionsResultAssertions]) *TypeClassicConditionsBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *TypeClassicConditionsBuilder) TimeRange(timeRange cog.Builder[expr.ExprTypeClassicConditionsTimeRange]) *TypeClassicConditionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TypeClassicConditions](./object-TypeClassicConditions.md)
