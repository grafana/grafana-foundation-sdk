---
title: <span class="badge builder"></span> DataQueryKindBuilder
---
# <span class="badge builder"></span> DataQueryKindBuilder

## Constructor

```go
func NewDataQueryKindBuilder() *DataQueryKindBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataQueryKindBuilder) Build() (DataQueryKind, error)
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *DataQueryKindBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *DataQueryKindBuilder
```

### <span class="badge object-method"></span> Group

```go
func (builder *DataQueryKindBuilder) Group(group string) *DataQueryKindBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *DataQueryKindBuilder) Spec(spec any) *DataQueryKindBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *DataQueryKindBuilder) Version(version string) *DataQueryKindBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DataQueryKind](./object-DataQueryKind.md)
