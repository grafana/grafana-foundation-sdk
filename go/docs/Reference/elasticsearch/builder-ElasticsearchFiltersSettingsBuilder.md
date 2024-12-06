---
title: <span class="badge builder"></span> ElasticsearchFiltersSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchFiltersSettingsBuilder

## Constructor

```go
func NewElasticsearchFiltersSettingsBuilder() *ElasticsearchFiltersSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchFiltersSettingsBuilder) Build() (ElasticsearchFiltersSettings, error)
```

### <span class="badge object-method"></span> Filters

```go
func (builder *ElasticsearchFiltersSettingsBuilder) Filters(filters []cog.Builder[elasticsearch.Filter]) *ElasticsearchFiltersSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchFiltersSettings](./object-ElasticsearchFiltersSettings.md)
