---
title: <span class="badge builder"></span> V2QueryGroupSpecBuilder
---
# <span class="badge builder"></span> V2QueryGroupSpecBuilder

## Constructor

```go
func NewV2QueryGroupSpecBuilder() *V2QueryGroupSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *V2QueryGroupSpecBuilder) Build() (V2QueryGroupSpec, error)
```

### <span class="badge object-method"></span> Queries

```go
func (builder *V2QueryGroupSpecBuilder) Queries(queries []cog.Builder[dashboardv2beta1.PanelQueryKind]) *V2QueryGroupSpecBuilder
```

### <span class="badge object-method"></span> QueryOptions

```go
func (builder *V2QueryGroupSpecBuilder) QueryOptions(queryOptions cog.Builder[dashboardv2beta1.QueryOptionsSpec]) *V2QueryGroupSpecBuilder
```

### <span class="badge object-method"></span> Transformations

```go
func (builder *V2QueryGroupSpecBuilder) Transformations(transformations []cog.Builder[dashboardv2beta1.V2TransformationKind]) *V2QueryGroupSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [V2QueryGroupSpec](./object-V2QueryGroupSpec.md)
