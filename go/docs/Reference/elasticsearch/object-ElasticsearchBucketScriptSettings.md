---
title: <span class="badge object-type-struct"></span> ElasticsearchBucketScriptSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchBucketScriptSettings

## Definition

```go
type ElasticsearchBucketScriptSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchBucketScriptSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchBucketScriptSettings *ElasticsearchBucketScriptSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchBucketScriptSettings` objects.

```go
func (elasticsearchBucketScriptSettings *ElasticsearchBucketScriptSettings) Equals(other ElasticsearchBucketScriptSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchBucketScriptSettings` fields for violations and returns them.

```go
func (elasticsearchBucketScriptSettings *ElasticsearchBucketScriptSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchBucketScriptSettingsBuilder](./builder-ElasticsearchBucketScriptSettingsBuilder.md)
