---
title: <span class="badge object-type-struct"></span> ElasticsearchSerialDiffSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchSerialDiffSettings

## Definition

```go
type ElasticsearchSerialDiffSettings struct {
    Lag *string `json:"lag,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchSerialDiffSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchSerialDiffSettings *ElasticsearchSerialDiffSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchSerialDiffSettings` objects.

```go
func (elasticsearchSerialDiffSettings *ElasticsearchSerialDiffSettings) Equals(other ElasticsearchSerialDiffSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchSerialDiffSettings` fields for violations and returns them.

```go
func (elasticsearchSerialDiffSettings *ElasticsearchSerialDiffSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchSerialDiffSettingsBuilder](./builder-ElasticsearchSerialDiffSettingsBuilder.md)
