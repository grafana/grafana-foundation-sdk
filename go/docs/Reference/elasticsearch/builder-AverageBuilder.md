---
title: <span class="badge builder"></span> AverageBuilder
---
# <span class="badge builder"></span> AverageBuilder

## Constructor

```go
func NewAverageBuilder() *AverageBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AverageBuilder) Build() (Average, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *AverageBuilder) Field(field string) *AverageBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *AverageBuilder) Hide(hide bool) *AverageBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *AverageBuilder) Id(id string) *AverageBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *AverageBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchAverageSettings]) *AverageBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Average](./object-Average.md)
