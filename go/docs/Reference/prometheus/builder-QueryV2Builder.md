---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```go
func NewQueryV2Builder() *QueryV2Builder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error)
```

### <span class="badge object-method"></span> AdhocFilters

Additional Ad-hoc filters that take precedence over Scope on conflict.

```go
func (builder *QueryV2Builder) AdhocFilters(adhocFilters []cog.Builder[prometheus.AdhocFilters]) *QueryV2Builder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> EditorMode

what we should show in the editor

Possible enum values:

 - `"builder"` 

 - `"code"` 

```go
func (builder *QueryV2Builder) EditorMode(editorMode prometheus.QueryEditorMode) *QueryV2Builder
```

### <span class="badge object-method"></span> Exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```go
func (builder *QueryV2Builder) Exemplar(exemplar bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Expr

The actual expression/query that will be evaluated by Prometheus

```go
func (builder *QueryV2Builder) Expr(expr string) *QueryV2Builder
```

### <span class="badge object-method"></span> Format

The response format

Possible enum values:

 - `"time_series"` 

 - `"table"` 

 - `"heatmap"` 

```go
func (builder *QueryV2Builder) Format(format prometheus.PromQueryFormat) *QueryV2Builder
```

### <span class="badge object-method"></span> GroupByKeys

Group By parameters to apply to aggregate expressions in the query

```go
func (builder *QueryV2Builder) GroupByKeys(groupByKeys []string) *QueryV2Builder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Instant

Returns only the latest value that Prometheus has scraped for the requested time series

```go
func (builder *QueryV2Builder) Instant(instant bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```go
func (builder *QueryV2Builder) Interval(interval string) *QueryV2Builder
```

### <span class="badge object-method"></span> IntervalFactor

Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

Deprecated: use interval

```go
func (builder *QueryV2Builder) IntervalFactor(intervalFactor int64) *QueryV2Builder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *QueryV2Builder) IntervalMs(intervalMs float64) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> LegendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```go
func (builder *QueryV2Builder) LegendFormat(legendFormat string) *QueryV2Builder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *QueryV2Builder) MaxDataPoints(maxDataPoints int64) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
```

### <span class="badge object-method"></span> Range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```go
func (builder *QueryV2Builder) Range(rangeArg bool) *QueryV2Builder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *QueryV2Builder) ResultAssertions(resultAssertions cog.Builder[prometheus.ResultAssertions]) *QueryV2Builder
```

### <span class="badge object-method"></span> Scopes

A set of filters applied to apply to the query

```go
func (builder *QueryV2Builder) Scopes(scopes []cog.Builder[prometheus.Scopes]) *QueryV2Builder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *QueryV2Builder) TimeRange(timeRange cog.Builder[prometheus.TimeRange]) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
