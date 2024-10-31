---
title: <span class="badge object-type-struct"></span> ScaleDistributionConfig
---
# <span class="badge object-type-struct"></span> ScaleDistributionConfig

TODO docs

## Definition

```go
type ScaleDistributionConfig struct {
    Type common.ScaleDistribution `json:"type"`
    Log *float64 `json:"log,omitempty"`
    LinearThreshold *float64 `json:"linearThreshold,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScaleDistributionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (scaleDistributionConfig *ScaleDistributionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ScaleDistributionConfig` objects.

```go
func (scaleDistributionConfig *ScaleDistributionConfig) Equals(other ScaleDistributionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ScaleDistributionConfig` fields for violations and returns them.

```go
func (scaleDistributionConfig *ScaleDistributionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ScaleDistributionConfigBuilder](./builder-ScaleDistributionConfigBuilder.md)
