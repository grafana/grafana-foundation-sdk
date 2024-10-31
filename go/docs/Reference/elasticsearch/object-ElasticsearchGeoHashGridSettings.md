---
title: <span class="badge object-type-struct"></span> ElasticsearchGeoHashGridSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchGeoHashGridSettings

## Definition

```go
type ElasticsearchGeoHashGridSettings struct {
    Precision *string `json:"precision,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchGeoHashGridSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchGeoHashGridSettings *ElasticsearchGeoHashGridSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchGeoHashGridSettings` objects.

```go
func (elasticsearchGeoHashGridSettings *ElasticsearchGeoHashGridSettings) Equals(other ElasticsearchGeoHashGridSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchGeoHashGridSettings` fields for violations and returns them.

```go
func (elasticsearchGeoHashGridSettings *ElasticsearchGeoHashGridSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchGeoHashGridSettingsBuilder](./builder-ElasticsearchGeoHashGridSettingsBuilder.md)
