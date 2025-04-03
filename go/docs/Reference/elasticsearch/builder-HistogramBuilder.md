---
title: <span class="badge builder"></span> HistogramBuilder
---
# <span class="badge builder"></span> HistogramBuilder

## Constructor

```go
func NewHistogramBuilder() *HistogramBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HistogramBuilder) Build() (Histogram, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *HistogramBuilder) Field(field string) *HistogramBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *HistogramBuilder) Id(id string) *HistogramBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *HistogramBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchHistogramSettings]) *HistogramBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Histogram](./object-Histogram.md)
