---
title: <span class="badge builder"></span> ElasticsearchMinSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchMinSettingsBuilder

## Constructor

```go
func NewElasticsearchMinSettingsBuilder() *ElasticsearchMinSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchMinSettingsBuilder) Build() (ElasticsearchMinSettings, error)
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchMinSettingsBuilder) Missing(missing string) *ElasticsearchMinSettingsBuilder
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchMinSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchMinSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchMinSettings](./object-ElasticsearchMinSettings.md)
