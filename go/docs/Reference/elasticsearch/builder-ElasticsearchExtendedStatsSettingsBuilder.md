---
title: <span class="badge builder"></span> ElasticsearchExtendedStatsSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchExtendedStatsSettingsBuilder

## Constructor

```go
func NewElasticsearchExtendedStatsSettingsBuilder() *ElasticsearchExtendedStatsSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchExtendedStatsSettingsBuilder) Build() (ElasticsearchExtendedStatsSettings, error)
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchExtendedStatsSettingsBuilder) Missing(missing string) *ElasticsearchExtendedStatsSettingsBuilder
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchExtendedStatsSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchExtendedStatsSettingsBuilder
```

### <span class="badge object-method"></span> Sigma

```go
func (builder *ElasticsearchExtendedStatsSettingsBuilder) Sigma(sigma string) *ElasticsearchExtendedStatsSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchExtendedStatsSettings](./object-ElasticsearchExtendedStatsSettings.md)
