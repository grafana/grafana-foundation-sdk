---
title: <span class="badge builder"></span> SumBuilder
---
# <span class="badge builder"></span> SumBuilder

## Constructor

```go
func NewSumBuilder() *SumBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SumBuilder) Build() (Sum, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *SumBuilder) Field(field string) *SumBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *SumBuilder) Hide(hide bool) *SumBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *SumBuilder) Id(id string) *SumBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *SumBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchSumSettings]) *SumBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Sum](./object-Sum.md)
