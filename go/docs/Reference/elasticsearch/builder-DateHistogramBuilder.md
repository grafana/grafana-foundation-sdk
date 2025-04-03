---
title: <span class="badge builder"></span> DateHistogramBuilder
---
# <span class="badge builder"></span> DateHistogramBuilder

## Constructor

```go
func NewDateHistogramBuilder() *DateHistogramBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DateHistogramBuilder) Build() (DateHistogram, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *DateHistogramBuilder) Field(field string) *DateHistogramBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *DateHistogramBuilder) Id(id string) *DateHistogramBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *DateHistogramBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchDateHistogramSettings]) *DateHistogramBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DateHistogram](./object-DateHistogram.md)
