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

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> EditorMode

Specifies which editor is being used to prepare the query. It can be "code" or "builder"

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

Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"

```go
func (builder *QueryV2Builder) Format(format prometheus.PromQueryFormat) *QueryV2Builder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

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

@deprecated Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

```go
func (builder *QueryV2Builder) IntervalFactor(intervalFactor float64) *QueryV2Builder
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

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
```

### <span class="badge object-method"></span> Range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```go
func (builder *QueryV2Builder) Range(rangeArg bool) *QueryV2Builder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
