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

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> PanelId

Panel ID from wich the queries will be reused.

```go
func (builder *QueryV2Builder) PanelId(panelId uint32) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
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

### <span class="badge object-method"></span> WithTransforms

```go
func (builder *QueryV2Builder) WithTransforms(withTransforms bool) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
