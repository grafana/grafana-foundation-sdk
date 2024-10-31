---
title: <span class="badge object-type-struct"></span> ElasticsearchAverageSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchAverageSettings

## Definition

```go
type ElasticsearchAverageSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchAverageSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchAverageSettings *ElasticsearchAverageSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchAverageSettings` objects.

```go
func (elasticsearchAverageSettings *ElasticsearchAverageSettings) Equals(other ElasticsearchAverageSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchAverageSettings` fields for violations and returns them.

```go
func (elasticsearchAverageSettings *ElasticsearchAverageSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchAverageSettingsBuilder](./builder-ElasticsearchAverageSettingsBuilder.md)
