---
title: <span class="badge object-type-struct"></span> RawDocument
---
# <span class="badge object-type-struct"></span> RawDocument

## Definition

```go
type RawDocument struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchRawDocumentSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RawDocument` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rawDocument *RawDocument) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RawDocument` objects.

```go
func (rawDocument *RawDocument) Equals(other RawDocument) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RawDocument` fields for violations and returns them.

```go
func (rawDocument *RawDocument) Validate() error
```

## See also

 * <span class="badge builder"></span> [RawDocumentBuilder](./builder-RawDocumentBuilder.md)
