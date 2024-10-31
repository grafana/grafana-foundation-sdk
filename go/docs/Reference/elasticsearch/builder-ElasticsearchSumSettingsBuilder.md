---
title: <span class="badge builder"></span> ElasticsearchSumSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchSumSettingsBuilder

## Constructor

```go
func NewElasticsearchSumSettingsBuilder() *ElasticsearchSumSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchSumSettingsBuilder) Build() (ElasticsearchSumSettings, error)
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchSumSettingsBuilder) Missing(missing string) *ElasticsearchSumSettingsBuilder
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchSumSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchSumSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchSumSettings](./object-ElasticsearchSumSettings.md)
