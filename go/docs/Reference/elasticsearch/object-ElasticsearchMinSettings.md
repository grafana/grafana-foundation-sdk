---
title: <span class="badge object-type-struct"></span> ElasticsearchMinSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMinSettings

## Definition

```go
type ElasticsearchMinSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMinSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMinSettings *ElasticsearchMinSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMinSettings` objects.

```go
func (elasticsearchMinSettings *ElasticsearchMinSettings) Equals(other ElasticsearchMinSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMinSettings` fields for violations and returns them.

```go
func (elasticsearchMinSettings *ElasticsearchMinSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMinSettingsBuilder](./builder-ElasticsearchMinSettingsBuilder.md)
