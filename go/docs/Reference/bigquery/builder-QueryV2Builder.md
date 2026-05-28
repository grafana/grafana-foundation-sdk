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

### <span class="badge object-method"></span> ConvertToUTC

```go
func (builder *QueryV2Builder) ConvertToUTC(convertToUTC bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Dataset

```go
func (builder *QueryV2Builder) Dataset(dataset string) *QueryV2Builder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> EditorMode

```go
func (builder *QueryV2Builder) EditorMode(editorMode bigquery.EditorMode) *QueryV2Builder
```

### <span class="badge object-method"></span> Format

```go
func (builder *QueryV2Builder) Format(format bigquery.QueryFormat) *QueryV2Builder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> Location

```go
func (builder *QueryV2Builder) Location(location string) *QueryV2Builder
```

### <span class="badge object-method"></span> Partitioned

```go
func (builder *QueryV2Builder) Partitioned(partitioned bool) *QueryV2Builder
```

### <span class="badge object-method"></span> PartitionedField

```go
func (builder *QueryV2Builder) PartitionedField(partitionedField string) *QueryV2Builder
```

### <span class="badge object-method"></span> Project

```go
func (builder *QueryV2Builder) Project(project string) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryPriority

```go
func (builder *QueryV2Builder) QueryPriority(queryPriority bigquery.QueryPriority) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *QueryV2Builder) RawQuery(rawQuery bool) *QueryV2Builder
```

### <span class="badge object-method"></span> RawSql

```go
func (builder *QueryV2Builder) RawSql(rawSql string) *QueryV2Builder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder
```

### <span class="badge object-method"></span> Sharded

```go
func (builder *QueryV2Builder) Sharded(sharded bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Sql

```go
func (builder *QueryV2Builder) Sql(sql cog.Builder[bigquery.SQLExpression]) *QueryV2Builder
```

### <span class="badge object-method"></span> Table

```go
func (builder *QueryV2Builder) Table(table string) *QueryV2Builder
```

### <span class="badge object-method"></span> TimeShift

```go
func (builder *QueryV2Builder) TimeShift(timeShift string) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
