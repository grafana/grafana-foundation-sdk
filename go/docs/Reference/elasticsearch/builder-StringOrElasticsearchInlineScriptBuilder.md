---
title: <span class="badge builder"></span> StringOrElasticsearchInlineScriptBuilder
---
# <span class="badge builder"></span> StringOrElasticsearchInlineScriptBuilder

## Constructor

```go
func NewStringOrElasticsearchInlineScriptBuilder() *StringOrElasticsearchInlineScriptBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StringOrElasticsearchInlineScriptBuilder) Build() (StringOrElasticsearchInlineScript, error)
```

### <span class="badge object-method"></span> ElasticsearchInlineScript

```go
func (builder *StringOrElasticsearchInlineScriptBuilder) ElasticsearchInlineScript(elasticsearchInlineScript cog.Builder[elasticsearch.ElasticsearchInlineScript]) *StringOrElasticsearchInlineScriptBuilder
```

### <span class="badge object-method"></span> String

```go
func (builder *StringOrElasticsearchInlineScriptBuilder) String(stringArg string) *StringOrElasticsearchInlineScriptBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [StringOrElasticsearchInlineScript](./object-StringOrElasticsearchInlineScript.md)
