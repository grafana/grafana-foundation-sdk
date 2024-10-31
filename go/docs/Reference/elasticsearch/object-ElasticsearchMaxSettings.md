---
title: <span class="badge object-type-struct"></span> ElasticsearchMaxSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMaxSettings

## Definition

```go
type ElasticsearchMaxSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMaxSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMaxSettings *ElasticsearchMaxSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMaxSettings` objects.

```go
func (elasticsearchMaxSettings *ElasticsearchMaxSettings) Equals(other ElasticsearchMaxSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMaxSettings` fields for violations and returns them.

```go
func (elasticsearchMaxSettings *ElasticsearchMaxSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMaxSettingsBuilder](./builder-ElasticsearchMaxSettingsBuilder.md)
