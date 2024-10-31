---
title: <span class="badge builder"></span> FiltersBuilder
---
# <span class="badge builder"></span> FiltersBuilder

## Constructor

```go
func NewFiltersBuilder() *FiltersBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FiltersBuilder) Build() (Filters, error)
```

### <span class="badge object-method"></span> Id

```go
func (builder *FiltersBuilder) Id(id string) *FiltersBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *FiltersBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchFiltersSettings]) *FiltersBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Filters](./object-Filters.md)
