---
title: <span class="badge object-type-struct"></span> ElasticsearchRawDocumentSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchRawDocumentSettings

## Definition

```go
type ElasticsearchRawDocumentSettings struct {
    Size *string `json:"size,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchRawDocumentSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchRawDocumentSettings *ElasticsearchRawDocumentSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchRawDocumentSettings` objects.

```go
func (elasticsearchRawDocumentSettings *ElasticsearchRawDocumentSettings) Equals(other ElasticsearchRawDocumentSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchRawDocumentSettings` fields for violations and returns them.

```go
func (elasticsearchRawDocumentSettings *ElasticsearchRawDocumentSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchRawDocumentSettingsBuilder](./builder-ElasticsearchRawDocumentSettingsBuilder.md)
