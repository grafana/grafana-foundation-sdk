---
title: <span class="badge object-type-struct"></span> ElasticsearchInlineScript
---
# <span class="badge object-type-struct"></span> ElasticsearchInlineScript

## Definition

```go
type ElasticsearchInlineScript struct {
    Inline *string `json:"inline,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchInlineScript` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchInlineScript *ElasticsearchInlineScript) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchInlineScript` objects.

```go
func (elasticsearchInlineScript *ElasticsearchInlineScript) Equals(other ElasticsearchInlineScript) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchInlineScript` fields for violations and returns them.

```go
func (elasticsearchInlineScript *ElasticsearchInlineScript) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchInlineScriptBuilder](./builder-ElasticsearchInlineScriptBuilder.md)
