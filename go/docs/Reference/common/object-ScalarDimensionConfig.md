---
title: <span class="badge object-type-struct"></span> ScalarDimensionConfig
---
# <span class="badge object-type-struct"></span> ScalarDimensionConfig

## Definition

```go
type ScalarDimensionConfig struct {
    Min float64 `json:"min"`
    Max float64 `json:"max"`
    Fixed *float64 `json:"fixed,omitempty"`
    // fixed: T -- will be added by each element
    Field *string `json:"field,omitempty"`
    Mode *common.ScalarDimensionMode `json:"mode,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScalarDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (scalarDimensionConfig *ScalarDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ScalarDimensionConfig` objects.

```go
func (scalarDimensionConfig *ScalarDimensionConfig) Equals(other ScalarDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ScalarDimensionConfig` fields for violations and returns them.

```go
func (scalarDimensionConfig *ScalarDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ScalarDimensionConfigBuilder](./builder-ScalarDimensionConfigBuilder.md)
