---
title: <span class="badge builder"></span> ElasticsearchMaxSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchMaxSettingsBuilder

## Constructor

```go
func NewElasticsearchMaxSettingsBuilder() *ElasticsearchMaxSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchMaxSettingsBuilder) Build() (ElasticsearchMaxSettings, error)
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchMaxSettingsBuilder) Missing(missing string) *ElasticsearchMaxSettingsBuilder
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchMaxSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchMaxSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchMaxSettings](./object-ElasticsearchMaxSettings.md)
