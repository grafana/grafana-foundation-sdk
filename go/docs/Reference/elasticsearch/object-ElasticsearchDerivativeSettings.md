---
title: <span class="badge object-type-struct"></span> ElasticsearchDerivativeSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchDerivativeSettings

## Definition

```go
type ElasticsearchDerivativeSettings struct {
    Unit *string `json:"unit,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchDerivativeSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchDerivativeSettings *ElasticsearchDerivativeSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchDerivativeSettings` objects.

```go
func (elasticsearchDerivativeSettings *ElasticsearchDerivativeSettings) Equals(other ElasticsearchDerivativeSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchDerivativeSettings` fields for violations and returns them.

```go
func (elasticsearchDerivativeSettings *ElasticsearchDerivativeSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchDerivativeSettingsBuilder](./builder-ElasticsearchDerivativeSettingsBuilder.md)
