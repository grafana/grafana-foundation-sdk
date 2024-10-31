---
title: <span class="badge builder"></span> ElasticsearchPercentilesSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchPercentilesSettingsBuilder

## Constructor

```go
func NewElasticsearchPercentilesSettingsBuilder() *ElasticsearchPercentilesSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchPercentilesSettingsBuilder) Build() (ElasticsearchPercentilesSettings, error)
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchPercentilesSettingsBuilder) Missing(missing string) *ElasticsearchPercentilesSettingsBuilder
```

### <span class="badge object-method"></span> Percents

```go
func (builder *ElasticsearchPercentilesSettingsBuilder) Percents(percents []string) *ElasticsearchPercentilesSettingsBuilder
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchPercentilesSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchPercentilesSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchPercentilesSettings](./object-ElasticsearchPercentilesSettings.md)
