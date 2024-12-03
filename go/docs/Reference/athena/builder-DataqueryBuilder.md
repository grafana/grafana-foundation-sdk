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

### <span class="badge object-method"></span> Column

```go
func (builder *DataqueryBuilder) Column(column string) *DataqueryBuilder
```

### <span class="badge object-method"></span> ConnectionArgs

```go
func (builder *DataqueryBuilder) ConnectionArgs(connectionArgs cog.Builder[athena.ConnectionArgs]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder
```

### <span class="badge object-method"></span> Format

```go
func (builder *DataqueryBuilder) Format(format athena.FormatOptions) *DataqueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryID

```go
func (builder *DataqueryBuilder) QueryID(queryID string) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RawSQL

```go
func (builder *DataqueryBuilder) RawSQL(rawSQL string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Table

```go
func (builder *DataqueryBuilder) Table(table string) *DataqueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
