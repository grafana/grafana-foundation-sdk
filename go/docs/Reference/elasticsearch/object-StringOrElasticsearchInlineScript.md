---
title: <span class="badge object-type-struct"></span> StringOrElasticsearchInlineScript
---
# <span class="badge object-type-struct"></span> StringOrElasticsearchInlineScript

## Definition

```go
type StringOrElasticsearchInlineScript struct {
    String *string `json:"String,omitempty"`
    ElasticsearchInlineScript *elasticsearch.ElasticsearchInlineScript `json:"ElasticsearchInlineScript,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrElasticsearchInlineScript` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrElasticsearchInlineScript *StringOrElasticsearchInlineScript) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrElasticsearchInlineScript` objects.

```go
func (stringOrElasticsearchInlineScript *StringOrElasticsearchInlineScript) Equals(other StringOrElasticsearchInlineScript) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrElasticsearchInlineScript` fields for violations and returns them.

```go
func (stringOrElasticsearchInlineScript *StringOrElasticsearchInlineScript) Validate() error
```

## See also

 * <span class="badge builder"></span> [StringOrElasticsearchInlineScriptBuilder](./builder-StringOrElasticsearchInlineScriptBuilder.md)
