---
title: <span class="badge object-type-struct"></span> AxisConfig
---
# <span class="badge object-type-struct"></span> AxisConfig

TODO docs

## Definition

```go
type AxisConfig struct {
    AxisPlacement *common.AxisPlacement `json:"axisPlacement,omitempty"`
    AxisColorMode *common.AxisColorMode `json:"axisColorMode,omitempty"`
    AxisLabel *string `json:"axisLabel,omitempty"`
    AxisWidth *float64 `json:"axisWidth,omitempty"`
    AxisSoftMin *float64 `json:"axisSoftMin,omitempty"`
    AxisSoftMax *float64 `json:"axisSoftMax,omitempty"`
    AxisGridShow *bool `json:"axisGridShow,omitempty"`
    ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
    AxisCenteredZero *bool `json:"axisCenteredZero,omitempty"`
    AxisBorderShow *bool `json:"axisBorderShow,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AxisConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (axisConfig *AxisConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AxisConfig` objects.

```go
func (axisConfig *AxisConfig) Equals(other AxisConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AxisConfig` fields for violations and returns them.

```go
func (axisConfig *AxisConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [AxisConfigBuilder](./builder-AxisConfigBuilder.md)
