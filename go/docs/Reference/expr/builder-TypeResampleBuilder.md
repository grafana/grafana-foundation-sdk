---
title: <span class="badge builder"></span> TypeResampleBuilder
---
# <span class="badge builder"></span> TypeResampleBuilder

## Constructor

```go
func NewTypeResampleBuilder() *TypeResampleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TypeResampleBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Datasource

The datasource

```go
func (builder *TypeResampleBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeResampleBuilder
```

### <span class="badge object-method"></span> Downsampler

The downsample function

Possible enum values:

 - `"sum"` 

 - `"mean"` 

 - `"min"` 

 - `"max"` 

 - `"count"` 

 - `"last"` 

 - `"median"` 

```go
func (builder *TypeResampleBuilder) Downsampler(downsampler expr.TypeResampleDownsampler) *TypeResampleBuilder
```

### <span class="badge object-method"></span> Expression

The math expression

```go
func (builder *TypeResampleBuilder) Expression(expression string) *TypeResampleBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *TypeResampleBuilder) Hide(hide bool) *TypeResampleBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *TypeResampleBuilder) IntervalMs(intervalMs float64) *TypeResampleBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *TypeResampleBuilder) MaxDataPoints(maxDataPoints int64) *TypeResampleBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *TypeResampleBuilder) QueryType(queryType string) *TypeResampleBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *TypeResampleBuilder) RefId(refId string) *TypeResampleBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *TypeResampleBuilder) ResultAssertions(resultAssertions cog.Builder[expr.ExprTypeResampleResultAssertions]) *TypeResampleBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *TypeResampleBuilder) TimeRange(timeRange cog.Builder[expr.ExprTypeResampleTimeRange]) *TypeResampleBuilder
```

### <span class="badge object-method"></span> Upsampler

The upsample function

Possible enum values:

 - `"pad"` Use the last seen value

 - `"backfilling"` backfill

 - `"fillna"` Do not fill values (nill)

```go
func (builder *TypeResampleBuilder) Upsampler(upsampler expr.TypeResampleUpsampler) *TypeResampleBuilder
```

### <span class="badge object-method"></span> Window

The time duration

```go
func (builder *TypeResampleBuilder) Window(window string) *TypeResampleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TypeResample](./object-TypeResample.md)
