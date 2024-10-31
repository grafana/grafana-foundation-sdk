---
title: <span class="badge object-type-struct"></span> TextDimensionConfig
---
# <span class="badge object-type-struct"></span> TextDimensionConfig

## Definition

```go
type TextDimensionConfig struct {
    Mode common.TextDimensionMode `json:"mode"`
    // fixed: T -- will be added by each element
    Field *string `json:"field,omitempty"`
    Fixed *string `json:"fixed,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TextDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (textDimensionConfig *TextDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TextDimensionConfig` objects.

```go
func (textDimensionConfig *TextDimensionConfig) Equals(other TextDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TextDimensionConfig` fields for violations and returns them.

```go
func (textDimensionConfig *TextDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [TextDimensionConfigBuilder](./builder-TextDimensionConfigBuilder.md)
