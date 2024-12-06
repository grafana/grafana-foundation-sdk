---
title: <span class="badge object-type-struct"></span> ElasticsearchTermsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchTermsSettings

## Definition

```go
type ElasticsearchTermsSettings struct {
    Order *elasticsearch.TermsOrder `json:"order,omitempty"`
    Size *string `json:"size,omitempty"`
    MinDocCount *string `json:"min_doc_count,omitempty"`
    OrderBy *string `json:"orderBy,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchTermsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchTermsSettings *ElasticsearchTermsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchTermsSettings` objects.

```go
func (elasticsearchTermsSettings *ElasticsearchTermsSettings) Equals(other ElasticsearchTermsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchTermsSettings` fields for violations and returns them.

```go
func (elasticsearchTermsSettings *ElasticsearchTermsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchTermsSettingsBuilder](./builder-ElasticsearchTermsSettingsBuilder.md)
