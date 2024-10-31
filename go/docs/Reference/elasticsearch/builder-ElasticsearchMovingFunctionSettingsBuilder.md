---
title: <span class="badge builder"></span> ElasticsearchMovingFunctionSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchMovingFunctionSettingsBuilder

## Constructor

```go
func NewElasticsearchMovingFunctionSettingsBuilder() *ElasticsearchMovingFunctionSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchMovingFunctionSettingsBuilder) Build() (ElasticsearchMovingFunctionSettings, error)
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchMovingFunctionSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchMovingFunctionSettingsBuilder
```

### <span class="badge object-method"></span> Shift

```go
func (builder *ElasticsearchMovingFunctionSettingsBuilder) Shift(shift string) *ElasticsearchMovingFunctionSettingsBuilder
```

### <span class="badge object-method"></span> Window

```go
func (builder *ElasticsearchMovingFunctionSettingsBuilder) Window(window string) *ElasticsearchMovingFunctionSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchMovingFunctionSettings](./object-ElasticsearchMovingFunctionSettings.md)
