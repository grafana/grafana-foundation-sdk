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

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> Direction

```go
func (builder *QueryBuilder) Direction(direction loki.LokiQueryDirection) *QueryBuilder
```

### <span class="badge object-method"></span> EditorMode

```go
func (builder *QueryBuilder) EditorMode(editorMode loki.QueryEditorMode) *QueryBuilder
```

### <span class="badge object-method"></span> Expr

The LogQL query.

```go
func (builder *QueryBuilder) Expr(expr string) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> Instant

@deprecated, now use queryType.

```go
func (builder *QueryBuilder) Instant(instant bool) *QueryBuilder
```

### <span class="badge object-method"></span> LegendFormat

Used to override the name of the series.

```go
func (builder *QueryBuilder) LegendFormat(legendFormat string) *QueryBuilder
```

### <span class="badge object-method"></span> MaxLines

Used to limit the number of log rows returned.

```go
func (builder *QueryBuilder) MaxLines(maxLines int64) *QueryBuilder
```

### <span class="badge object-method"></span> OldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> Range

@deprecated, now use queryType.

```go
func (builder *QueryBuilder) Range(rangeArg bool) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> Resolution

@deprecated, now use step.

```go
func (builder *QueryBuilder) Resolution(resolution int64) *QueryBuilder
```

### <span class="badge object-method"></span> Step

Used to set step value for range queries.

```go
func (builder *QueryBuilder) Step(step string) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
