---
title: <span class="badge builder"></span> MinBuilder
---
# <span class="badge builder"></span> MinBuilder

## Constructor

```go
func NewMinBuilder() *MinBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MinBuilder) Build() (Min, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *MinBuilder) Field(field string) *MinBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MinBuilder) Hide(hide bool) *MinBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MinBuilder) Id(id string) *MinBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MinBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMinSettings]) *MinBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Min](./object-Min.md)
