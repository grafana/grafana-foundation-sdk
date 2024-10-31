---
title: <span class="badge object-type-struct"></span> ElasticsearchMovingFunctionSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMovingFunctionSettings

## Definition

```go
type ElasticsearchMovingFunctionSettings struct {
    Window *string `json:"window,omitempty"`
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
    Shift *string `json:"shift,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingFunctionSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMovingFunctionSettings *ElasticsearchMovingFunctionSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMovingFunctionSettings` objects.

```go
func (elasticsearchMovingFunctionSettings *ElasticsearchMovingFunctionSettings) Equals(other ElasticsearchMovingFunctionSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMovingFunctionSettings` fields for violations and returns them.

```go
func (elasticsearchMovingFunctionSettings *ElasticsearchMovingFunctionSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMovingFunctionSettingsBuilder](./builder-ElasticsearchMovingFunctionSettingsBuilder.md)
