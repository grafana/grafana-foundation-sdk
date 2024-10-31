---
title: <span class="badge object-type-struct"></span> ElasticsearchSumSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchSumSettings

## Definition

```go
type ElasticsearchSumSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchSumSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchSumSettings *ElasticsearchSumSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchSumSettings` objects.

```go
func (elasticsearchSumSettings *ElasticsearchSumSettings) Equals(other ElasticsearchSumSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchSumSettings` fields for violations and returns them.

```go
func (elasticsearchSumSettings *ElasticsearchSumSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchSumSettingsBuilder](./builder-ElasticsearchSumSettingsBuilder.md)
