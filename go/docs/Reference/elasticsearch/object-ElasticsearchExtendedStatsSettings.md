---
title: <span class="badge object-type-struct"></span> ElasticsearchExtendedStatsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchExtendedStatsSettings

## Definition

```go
type ElasticsearchExtendedStatsSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Missing *string `json:"missing,omitempty"`
    Sigma *string `json:"sigma,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchExtendedStatsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchExtendedStatsSettings *ElasticsearchExtendedStatsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchExtendedStatsSettings` objects.

```go
func (elasticsearchExtendedStatsSettings *ElasticsearchExtendedStatsSettings) Equals(other ElasticsearchExtendedStatsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchExtendedStatsSettings` fields for violations and returns them.

```go
func (elasticsearchExtendedStatsSettings *ElasticsearchExtendedStatsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchExtendedStatsSettingsBuilder](./builder-ElasticsearchExtendedStatsSettingsBuilder.md)
