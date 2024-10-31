---
title: <span class="badge object-type-struct"></span> ElasticsearchRawDataSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchRawDataSettings

## Definition

```go
type ElasticsearchRawDataSettings struct {
    Size *string `json:"size,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchRawDataSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchRawDataSettings *ElasticsearchRawDataSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchRawDataSettings` objects.

```go
func (elasticsearchRawDataSettings *ElasticsearchRawDataSettings) Equals(other ElasticsearchRawDataSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchRawDataSettings` fields for violations and returns them.

```go
func (elasticsearchRawDataSettings *ElasticsearchRawDataSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchRawDataSettingsBuilder](./builder-ElasticsearchRawDataSettingsBuilder.md)
