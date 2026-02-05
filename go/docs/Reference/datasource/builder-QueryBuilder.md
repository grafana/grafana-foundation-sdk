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

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> PanelId

Panel ID from wich the queries will be reused.

```go
func (builder *QueryBuilder) PanelId(panelId uint32) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

### <span class="badge object-method"></span> WithTransforms

```go
func (builder *QueryBuilder) WithTransforms(withTransforms bool) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
