---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```go
func NewQueryBuilder(refId string) *QueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryBuilder) Build() (Query, error)
```

### <span class="badge object-method"></span> DatasourceUid

Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.

```go
func (builder *QueryBuilder) DatasourceUid(datasourceUid string) *QueryBuilder
```

### <span class="badge object-method"></span> Model

JSON is the raw JSON query and includes the above properties as well as custom properties.

```go
func (builder *QueryBuilder) Model(model cog.Builder[cog/variants.Dataquery]) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> RelativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

```go
func (builder *QueryBuilder) RelativeTimeRange(from alerting.Duration, to alerting.Duration) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Query](./object-Query.md)
