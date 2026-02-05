---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```go
func NewDataqueryBuilder() *DataqueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataqueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> ConvertToUTC

```go
func (builder *DataqueryBuilder) ConvertToUTC(convertToUTC bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Dataset

```go
func (builder *DataqueryBuilder) Dataset(dataset string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *DataqueryBuilder) Datasource(datasource common.DataSourceRef) *DataqueryBuilder
```

### <span class="badge object-method"></span> EditorMode

```go
func (builder *DataqueryBuilder) EditorMode(editorMode bigquery.EditorMode) *DataqueryBuilder
```

### <span class="badge object-method"></span> Format

```go
func (builder *DataqueryBuilder) Format(format bigquery.QueryFormat) *DataqueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Location

```go
func (builder *DataqueryBuilder) Location(location string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Partitioned

```go
func (builder *DataqueryBuilder) Partitioned(partitioned bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> PartitionedField

```go
func (builder *DataqueryBuilder) PartitionedField(partitionedField string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Project

```go
func (builder *DataqueryBuilder) Project(project string) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryPriority

```go
func (builder *DataqueryBuilder) QueryPriority(queryPriority bigquery.QueryPriority) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *DataqueryBuilder) RawQuery(rawQuery bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> RawSql

```go
func (builder *DataqueryBuilder) RawSql(rawSql string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Sharded

```go
func (builder *DataqueryBuilder) Sharded(sharded bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Sql

```go
func (builder *DataqueryBuilder) Sql(sql cog.Builder[bigquery.SQLExpression]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Table

```go
func (builder *DataqueryBuilder) Table(table string) *DataqueryBuilder
```

### <span class="badge object-method"></span> TimeShift

```go
func (builder *DataqueryBuilder) TimeShift(timeShift string) *DataqueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
