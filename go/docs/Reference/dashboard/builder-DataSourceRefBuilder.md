---
title: <span class="badge builder"></span> DataSourceRefBuilder
---
# <span class="badge builder"></span> DataSourceRefBuilder

## Constructor

```go
func NewDataSourceRefBuilder() *DataSourceRefBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataSourceRefBuilder) Build() (DataSourceRef, error)
```

### <span class="badge object-method"></span> Type

The plugin type-id

```go
func (builder *DataSourceRefBuilder) Type(typeArg string) *DataSourceRefBuilder
```

### <span class="badge object-method"></span> Uid

Specific datasource instance

```go
func (builder *DataSourceRefBuilder) Uid(uid string) *DataSourceRefBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [DataSourceRef](./object-DataSourceRef.md)
