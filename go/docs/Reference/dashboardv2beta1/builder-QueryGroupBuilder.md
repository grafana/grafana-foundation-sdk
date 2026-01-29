---
title: <span class="badge builder"></span> QueryGroupBuilder
---
# <span class="badge builder"></span> QueryGroupBuilder

## Constructor

```go
func NewQueryGroupBuilder() *QueryGroupBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryGroupBuilder) Build() (QueryGroupKind, error)
```

### <span class="badge object-method"></span> QueryOptions

```go
func (builder *QueryGroupBuilder) QueryOptions(queryOptions cog.Builder[dashboardv2beta1.QueryOptionsSpec]) *QueryGroupBuilder
```

### <span class="badge object-method"></span> Target

```go
func (builder *QueryGroupBuilder) Target(querie cog.Builder[dashboardv2beta1.PanelQueryKind]) *QueryGroupBuilder
```

### <span class="badge object-method"></span> Targets

```go
func (builder *QueryGroupBuilder) Targets(queries []cog.Builder[dashboardv2beta1.PanelQueryKind]) *QueryGroupBuilder
```

### <span class="badge object-method"></span> Transformation

```go
func (builder *QueryGroupBuilder) Transformation(transformation cog.Builder[dashboardv2beta1.TransformationKind]) *QueryGroupBuilder
```

### <span class="badge object-method"></span> Transformations

```go
func (builder *QueryGroupBuilder) Transformations(transformations []cog.Builder[dashboardv2beta1.TransformationKind]) *QueryGroupBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryGroupKind](./object-QueryGroupKind.md)
