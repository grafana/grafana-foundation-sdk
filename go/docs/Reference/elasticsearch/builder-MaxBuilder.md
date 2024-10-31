---
title: <span class="badge builder"></span> MaxBuilder
---
# <span class="badge builder"></span> MaxBuilder

## Constructor

```go
func NewMaxBuilder() *MaxBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MaxBuilder) Build() (Max, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *MaxBuilder) Field(field string) *MaxBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MaxBuilder) Hide(hide bool) *MaxBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MaxBuilder) Id(id string) *MaxBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MaxBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMaxSettings]) *MaxBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Max](./object-Max.md)
