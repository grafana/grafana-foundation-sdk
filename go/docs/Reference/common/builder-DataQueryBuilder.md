---
title: <span class="badge builder"></span> DataQueryBuilder
---
# <span class="badge builder"></span> DataQueryBuilder

## Constructor

```go
func NewDataQueryBuilder() *DataQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataQueryBuilder) Build() (DataQuery, error)
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *DataQueryBuilder) Datasource(datasource any) *DataQueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *DataQueryBuilder) Hide(hide bool) *DataQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *DataQueryBuilder) QueryType(queryType string) *DataQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *DataQueryBuilder) RefId(refId string) *DataQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DataQuery](./object-DataQuery.md)
