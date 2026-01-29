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

### <span class="badge object-method"></span> TypeClassicConditions

```go
func (builder *QueryBuilder) TypeClassicConditions(typeClassicConditions cog.Builder[expr.TypeClassicConditions]) *QueryBuilder
```

### <span class="badge object-method"></span> TypeMath

```go
func (builder *QueryBuilder) TypeMath(typeMath cog.Builder[expr.TypeMath]) *QueryBuilder
```

### <span class="badge object-method"></span> TypeReduce

```go
func (builder *QueryBuilder) TypeReduce(typeReduce cog.Builder[expr.TypeReduce]) *QueryBuilder
```

### <span class="badge object-method"></span> TypeResample

```go
func (builder *QueryBuilder) TypeResample(typeResample cog.Builder[expr.TypeResample]) *QueryBuilder
```

### <span class="badge object-method"></span> TypeSql

```go
func (builder *QueryBuilder) TypeSql(typeSql cog.Builder[expr.TypeSql]) *QueryBuilder
```

### <span class="badge object-method"></span> TypeThreshold

```go
func (builder *QueryBuilder) TypeThreshold(typeThreshold cog.Builder[expr.TypeThreshold]) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
