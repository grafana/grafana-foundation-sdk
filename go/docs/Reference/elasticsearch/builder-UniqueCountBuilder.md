---
title: <span class="badge builder"></span> UniqueCountBuilder
---
# <span class="badge builder"></span> UniqueCountBuilder

## Constructor

```go
func NewUniqueCountBuilder() *UniqueCountBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *UniqueCountBuilder) Build() (UniqueCount, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *UniqueCountBuilder) Field(field string) *UniqueCountBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *UniqueCountBuilder) Hide(hide bool) *UniqueCountBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *UniqueCountBuilder) Id(id string) *UniqueCountBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *UniqueCountBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchUniqueCountSettings]) *UniqueCountBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [UniqueCount](./object-UniqueCount.md)
