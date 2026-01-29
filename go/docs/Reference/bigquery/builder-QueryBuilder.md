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

### <span class="badge object-method"></span> ConvertToUTC

```go
func (builder *QueryBuilder) ConvertToUTC(convertToUTC bool) *QueryBuilder
```

### <span class="badge object-method"></span> Dataset

```go
func (builder *QueryBuilder) Dataset(dataset string) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> EditorMode

```go
func (builder *QueryBuilder) EditorMode(editorMode bigquery.EditorMode) *QueryBuilder
```

### <span class="badge object-method"></span> Format

```go
func (builder *QueryBuilder) Format(format bigquery.QueryFormat) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> Location

```go
func (builder *QueryBuilder) Location(location string) *QueryBuilder
```

### <span class="badge object-method"></span> OldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder
```

### <span class="badge object-method"></span> Partitioned

```go
func (builder *QueryBuilder) Partitioned(partitioned bool) *QueryBuilder
```

### <span class="badge object-method"></span> PartitionedField

```go
func (builder *QueryBuilder) PartitionedField(partitionedField string) *QueryBuilder
```

### <span class="badge object-method"></span> Project

```go
func (builder *QueryBuilder) Project(project string) *QueryBuilder
```

### <span class="badge object-method"></span> QueryPriority

```go
func (builder *QueryBuilder) QueryPriority(queryPriority bigquery.QueryPriority) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *QueryBuilder) RawQuery(rawQuery bool) *QueryBuilder
```

### <span class="badge object-method"></span> RawSql

```go
func (builder *QueryBuilder) RawSql(rawSql string) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> Sharded

```go
func (builder *QueryBuilder) Sharded(sharded bool) *QueryBuilder
```

### <span class="badge object-method"></span> Sql

```go
func (builder *QueryBuilder) Sql(sql cog.Builder[bigquery.SQLExpression]) *QueryBuilder
```

### <span class="badge object-method"></span> Table

```go
func (builder *QueryBuilder) Table(table string) *QueryBuilder
```

### <span class="badge object-method"></span> TimeShift

```go
func (builder *QueryBuilder) TimeShift(timeShift string) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
