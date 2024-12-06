---
title: <span class="badge object-type-struct"></span> ElasticsearchPercentilesSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchPercentilesSettings

## Definition

```go
type ElasticsearchPercentilesSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Missing *string `json:"missing,omitempty"`
    Percents []string `json:"percents,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchPercentilesSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchPercentilesSettings *ElasticsearchPercentilesSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchPercentilesSettings` objects.

```go
func (elasticsearchPercentilesSettings *ElasticsearchPercentilesSettings) Equals(other ElasticsearchPercentilesSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchPercentilesSettings` fields for violations and returns them.

```go
func (elasticsearchPercentilesSettings *ElasticsearchPercentilesSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchPercentilesSettingsBuilder](./builder-ElasticsearchPercentilesSettingsBuilder.md)
