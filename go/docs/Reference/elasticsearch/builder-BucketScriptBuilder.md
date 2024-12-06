---
title: <span class="badge builder"></span> BucketScriptBuilder
---
# <span class="badge builder"></span> BucketScriptBuilder

## Constructor

```go
func NewBucketScriptBuilder() *BucketScriptBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BucketScriptBuilder) Build() (BucketScript, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *BucketScriptBuilder) Hide(hide bool) *BucketScriptBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *BucketScriptBuilder) Id(id string) *BucketScriptBuilder
```

### <span class="badge object-method"></span> PipelineVariables

```go
func (builder *BucketScriptBuilder) PipelineVariables(pipelineVariables []cog.Builder[elasticsearch.PipelineVariable]) *BucketScriptBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *BucketScriptBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchBucketScriptSettings]) *BucketScriptBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BucketScript](./object-BucketScript.md)
