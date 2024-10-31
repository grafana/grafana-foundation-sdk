---
title: <span class="badge builder"></span> PercentilesBuilder
---
# <span class="badge builder"></span> PercentilesBuilder

## Constructor

```go
func NewPercentilesBuilder() *PercentilesBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PercentilesBuilder) Build() (Percentiles, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *PercentilesBuilder) Field(field string) *PercentilesBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *PercentilesBuilder) Hide(hide bool) *PercentilesBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *PercentilesBuilder) Id(id string) *PercentilesBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *PercentilesBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchPercentilesSettings]) *PercentilesBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Percentiles](./object-Percentiles.md)
