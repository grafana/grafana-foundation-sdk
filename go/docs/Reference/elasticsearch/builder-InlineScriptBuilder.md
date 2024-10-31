---
title: <span class="badge builder"></span> InlineScriptBuilder
---
# <span class="badge builder"></span> InlineScriptBuilder

## Constructor

```go
func NewInlineScriptBuilder() *InlineScriptBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *InlineScriptBuilder) Build() (InlineScript, error)
```

### <span class="badge object-method"></span> ElasticsearchInlineScript

```go
func (builder *InlineScriptBuilder) ElasticsearchInlineScript(elasticsearchInlineScript cog.Builder[elasticsearch.ElasticsearchInlineScript]) *InlineScriptBuilder
```

### <span class="badge object-method"></span> String

```go
func (builder *InlineScriptBuilder) String(stringArg string) *InlineScriptBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [InlineScript](./object-InlineScript.md)
