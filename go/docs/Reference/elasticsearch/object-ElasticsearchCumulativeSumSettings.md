---
title: <span class="badge object-type-struct"></span> ElasticsearchCumulativeSumSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchCumulativeSumSettings

## Definition

```go
type ElasticsearchCumulativeSumSettings struct {
    Format *string `json:"format,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchCumulativeSumSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchCumulativeSumSettings *ElasticsearchCumulativeSumSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchCumulativeSumSettings` objects.

```go
func (elasticsearchCumulativeSumSettings *ElasticsearchCumulativeSumSettings) Equals(other ElasticsearchCumulativeSumSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchCumulativeSumSettings` fields for violations and returns them.

```go
func (elasticsearchCumulativeSumSettings *ElasticsearchCumulativeSumSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchCumulativeSumSettingsBuilder](./builder-ElasticsearchCumulativeSumSettingsBuilder.md)
