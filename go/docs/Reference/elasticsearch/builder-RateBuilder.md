---
title: <span class="badge builder"></span> RateBuilder
---
# <span class="badge builder"></span> RateBuilder

## Constructor

```go
func NewRateBuilder() *RateBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RateBuilder) Build() (Rate, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *RateBuilder) Field(field string) *RateBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *RateBuilder) Hide(hide bool) *RateBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *RateBuilder) Id(id string) *RateBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *RateBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchRateSettings]) *RateBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Rate](./object-Rate.md)
