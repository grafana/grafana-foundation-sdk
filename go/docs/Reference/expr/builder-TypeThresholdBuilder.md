---
title: <span class="badge builder"></span> TypeThresholdBuilder
---
# <span class="badge builder"></span> TypeThresholdBuilder

## Constructor

```go
func NewTypeThresholdBuilder() *TypeThresholdBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TypeThresholdBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Conditions

Threshold Conditions

```go
func (builder *TypeThresholdBuilder) Conditions(conditions []cog.Builder[expr.ExprTypeThresholdConditions]) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> Datasource

The datasource

```go
func (builder *TypeThresholdBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> Expression

Reference to single query result

```go
func (builder *TypeThresholdBuilder) Expression(expression string) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *TypeThresholdBuilder) Hide(hide bool) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *TypeThresholdBuilder) IntervalMs(intervalMs float64) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *TypeThresholdBuilder) MaxDataPoints(maxDataPoints int64) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *TypeThresholdBuilder) QueryType(queryType string) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *TypeThresholdBuilder) RefId(refId string) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *TypeThresholdBuilder) ResultAssertions(resultAssertions cog.Builder[expr.ExprTypeThresholdResultAssertions]) *TypeThresholdBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *TypeThresholdBuilder) TimeRange(timeRange cog.Builder[expr.ExprTypeThresholdTimeRange]) *TypeThresholdBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TypeThreshold](./object-TypeThreshold.md)
