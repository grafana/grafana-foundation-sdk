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

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder
```

### <span class="badge object-method"></span> GroupBy

Allows to group the results.

```go
func (builder *DataqueryBuilder) GroupBy(groupBy []string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> LabelSelector

Specifies the query label selectors.

```go
func (builder *DataqueryBuilder) LabelSelector(labelSelector string) *DataqueryBuilder
```

### <span class="badge object-method"></span> MaxNodes

Sets the maximum number of nodes in the flamegraph.

```go
func (builder *DataqueryBuilder) MaxNodes(maxNodes int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> ProfileTypeId

Specifies the type of profile to query.

```go
func (builder *DataqueryBuilder) ProfileTypeId(profileTypeId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> SpanSelector

Specifies the query span selectors.

```go
func (builder *DataqueryBuilder) SpanSelector(spanSelector []string) *DataqueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
