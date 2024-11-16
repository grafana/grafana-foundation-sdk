---
title: <span class="badge object-type-struct"></span> ElasticsearchUniqueCountSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchUniqueCountSettings

## Definition

```go
type ElasticsearchUniqueCountSettings struct {
    PrecisionThreshold *string `json:"precision_threshold,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchUniqueCountSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchUniqueCountSettings *ElasticsearchUniqueCountSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchUniqueCountSettings` objects.

```go
func (elasticsearchUniqueCountSettings *ElasticsearchUniqueCountSettings) Equals(other ElasticsearchUniqueCountSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchUniqueCountSettings` fields for violations and returns them.

```go
func (elasticsearchUniqueCountSettings *ElasticsearchUniqueCountSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchUniqueCountSettingsBuilder](./builder-ElasticsearchUniqueCountSettingsBuilder.md)