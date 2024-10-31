---
title: <span class="badge builder"></span> ElasticsearchBucketScriptSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchBucketScriptSettingsBuilder

## Constructor

```go
func NewElasticsearchBucketScriptSettingsBuilder() *ElasticsearchBucketScriptSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchBucketScriptSettingsBuilder) Build() (ElasticsearchBucketScriptSettings, error)
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchBucketScriptSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchBucketScriptSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchBucketScriptSettings](./object-ElasticsearchBucketScriptSettings.md)
