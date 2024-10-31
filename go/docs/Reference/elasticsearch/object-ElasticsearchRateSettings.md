---
title: <span class="badge object-type-struct"></span> ElasticsearchRateSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchRateSettings

## Definition

```go
type ElasticsearchRateSettings struct {
    Unit *string `json:"unit,omitempty"`
    Mode *string `json:"mode,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchRateSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchRateSettings *ElasticsearchRateSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchRateSettings` objects.

```go
func (elasticsearchRateSettings *ElasticsearchRateSettings) Equals(other ElasticsearchRateSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchRateSettings` fields for violations and returns them.

```go
func (elasticsearchRateSettings *ElasticsearchRateSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchRateSettingsBuilder](./builder-ElasticsearchRateSettingsBuilder.md)
