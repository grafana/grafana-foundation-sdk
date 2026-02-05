---
title: <span class="badge object-type-struct"></span> TermsSettings
---
# <span class="badge object-type-struct"></span> TermsSettings

## Definition

```go
type TermsSettings struct {
    Order *elasticsearch.TermsOrder `json:"order,omitempty"`
    Size *string `json:"size,omitempty"`
    MinDocCount *string `json:"min_doc_count,omitempty"`
    OrderBy *string `json:"orderBy,omitempty"`
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TermsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (termsSettings *TermsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TermsSettings` objects.

```go
func (termsSettings *TermsSettings) Equals(other TermsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TermsSettings` fields for violations and returns them.

```go
func (termsSettings *TermsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [TermsSettingsBuilder](./builder-TermsSettingsBuilder.md)
