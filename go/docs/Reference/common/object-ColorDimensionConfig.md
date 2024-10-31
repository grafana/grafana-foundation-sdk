---
title: <span class="badge object-type-struct"></span> ColorDimensionConfig
---
# <span class="badge object-type-struct"></span> ColorDimensionConfig

## Definition

```go
type ColorDimensionConfig struct {
    // color value
    Fixed *string `json:"fixed,omitempty"`
    // fixed: T -- will be added by each element
    Field *string `json:"field,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ColorDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (colorDimensionConfig *ColorDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ColorDimensionConfig` objects.

```go
func (colorDimensionConfig *ColorDimensionConfig) Equals(other ColorDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ColorDimensionConfig` fields for violations and returns them.

```go
func (colorDimensionConfig *ColorDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ColorDimensionConfigBuilder](./builder-ColorDimensionConfigBuilder.md)
