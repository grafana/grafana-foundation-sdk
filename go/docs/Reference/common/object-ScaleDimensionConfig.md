---
title: <span class="badge object-type-struct"></span> ScaleDimensionConfig
---
# <span class="badge object-type-struct"></span> ScaleDimensionConfig

## Definition

```go
type ScaleDimensionConfig struct {
    Min float64 `json:"min"`
    Max float64 `json:"max"`
    Fixed *float64 `json:"fixed,omitempty"`
    // fixed: T -- will be added by each element
    Field *string `json:"field,omitempty"`
    // | *"linear"
    Mode *common.ScaleDimensionMode `json:"mode,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScaleDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (scaleDimensionConfig *ScaleDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ScaleDimensionConfig` objects.

```go
func (scaleDimensionConfig *ScaleDimensionConfig) Equals(other ScaleDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ScaleDimensionConfig` fields for violations and returns them.

```go
func (scaleDimensionConfig *ScaleDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ScaleDimensionConfigBuilder](./builder-ScaleDimensionConfigBuilder.md)
