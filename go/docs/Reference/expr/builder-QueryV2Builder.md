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

### <span class="badge object-method"></span> TypeClassicConditions

```go
func (builder *QueryV2Builder) TypeClassicConditions(typeClassicConditions cog.Builder[expr.TypeClassicConditions]) *QueryV2Builder
```

### <span class="badge object-method"></span> TypeMath

```go
func (builder *QueryV2Builder) TypeMath(typeMath cog.Builder[expr.TypeMath]) *QueryV2Builder
```

### <span class="badge object-method"></span> TypeReduce

```go
func (builder *QueryV2Builder) TypeReduce(typeReduce cog.Builder[expr.TypeReduce]) *QueryV2Builder
```

### <span class="badge object-method"></span> TypeResample

```go
func (builder *QueryV2Builder) TypeResample(typeResample cog.Builder[expr.TypeResample]) *QueryV2Builder
```

### <span class="badge object-method"></span> TypeSql

```go
func (builder *QueryV2Builder) TypeSql(typeSql cog.Builder[expr.TypeSql]) *QueryV2Builder
```

### <span class="badge object-method"></span> TypeThreshold

```go
func (builder *QueryV2Builder) TypeThreshold(typeThreshold cog.Builder[expr.TypeThreshold]) *QueryV2Builder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
