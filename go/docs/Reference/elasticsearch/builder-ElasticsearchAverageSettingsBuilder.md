---
title: <span class="badge builder"></span> ElasticsearchAverageSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchAverageSettingsBuilder

## Constructor

```go
func NewElasticsearchAverageSettingsBuilder() *ElasticsearchAverageSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchAverageSettingsBuilder) Build() (ElasticsearchAverageSettings, error)
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchAverageSettingsBuilder) Missing(missing string) *ElasticsearchAverageSettingsBuilder
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchAverageSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchAverageSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchAverageSettings](./object-ElasticsearchAverageSettings.md)
