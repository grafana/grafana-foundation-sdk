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

### <span class="badge object-method"></span> AdhocFilters

Additional Ad-hoc filters that take precedence over Scope on conflict.

```go
func (builder *QueryBuilder) AdhocFilters(adhocFilters []cog.Builder[prometheus.AdhocFilters]) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> EditorMode

what we should show in the editor

Possible enum values:

 - `"builder"` 

 - `"code"` 

```go
func (builder *QueryBuilder) EditorMode(editorMode prometheus.QueryEditorMode) *QueryBuilder
```

### <span class="badge object-method"></span> Exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```go
func (builder *QueryBuilder) Exemplar(exemplar bool) *QueryBuilder
```

### <span class="badge object-method"></span> Expr

The actual expression/query that will be evaluated by Prometheus

```go
func (builder *QueryBuilder) Expr(expr string) *QueryBuilder
```

### <span class="badge object-method"></span> Format

The response format

Possible enum values:

 - `"time_series"` 

 - `"table"` 

 - `"heatmap"` 

```go
func (builder *QueryBuilder) Format(format prometheus.PromQueryFormat) *QueryBuilder
```

### <span class="badge object-method"></span> GroupByKeys

Group By parameters to apply to aggregate expressions in the query

```go
func (builder *QueryBuilder) GroupByKeys(groupByKeys []string) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> Instant

Returns only the latest value that Prometheus has scraped for the requested time series

```go
func (builder *QueryBuilder) Instant() *QueryBuilder
```

### <span class="badge object-method"></span> Interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```go
func (builder *QueryBuilder) Interval(interval string) *QueryBuilder
```

### <span class="badge object-method"></span> IntervalFactor

Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

Deprecated: use interval

```go
func (builder *QueryBuilder) IntervalFactor(intervalFactor int64) *QueryBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *QueryBuilder) IntervalMs(intervalMs float64) *QueryBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryBuilder) Labels(labels map[string]string) *QueryBuilder
```

### <span class="badge object-method"></span> LegendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```go
func (builder *QueryBuilder) LegendFormat(legendFormat string) *QueryBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *QueryBuilder) MaxDataPoints(maxDataPoints int64) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> Range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```go
func (builder *QueryBuilder) Range() *QueryBuilder
```

### <span class="badge object-method"></span> RangeAndInstant

```go
func (builder *QueryBuilder) RangeAndInstant() *QueryBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *QueryBuilder) ResultAssertions(resultAssertions cog.Builder[prometheus.ResultAssertions]) *QueryBuilder
```

### <span class="badge object-method"></span> Scopes

A set of filters applied to apply to the query

```go
func (builder *QueryBuilder) Scopes(scopes []cog.Builder[prometheus.Scopes]) *QueryBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *QueryBuilder) TimeRange(timeRange cog.Builder[prometheus.TimeRange]) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
