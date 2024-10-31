---
title: <span class="badge object-type-struct"></span> YAxisConfig
---
# <span class="badge object-type-struct"></span> YAxisConfig

Configuration options for the yAxis

## Definition

```go
type YAxisConfig struct {
    // Sets the yAxis unit
    Unit *string `json:"unit,omitempty"`
    // Reverses the yAxis
    Reverse *bool `json:"reverse,omitempty"`
    // Controls the number of decimals for yAxis values
    Decimals *float32 `json:"decimals,omitempty"`
    // Sets the minimum value for the yAxis
    Min *float32 `json:"min,omitempty"`
    AxisPlacement *common.AxisPlacement `json:"axisPlacement,omitempty"`
    AxisColorMode *common.AxisColorMode `json:"axisColorMode,omitempty"`
    AxisLabel *string `json:"axisLabel,omitempty"`
    AxisWidth *float64 `json:"axisWidth,omitempty"`
    AxisSoftMin *float64 `json:"axisSoftMin,omitempty"`
    AxisSoftMax *float64 `json:"axisSoftMax,omitempty"`
    AxisGridShow *bool `json:"axisGridShow,omitempty"`
    ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
    AxisCenteredZero *bool `json:"axisCenteredZero,omitempty"`
    // Sets the maximum value for the yAxis
    Max *float32 `json:"max,omitempty"`
    AxisBorderShow *bool `json:"axisBorderShow,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `YAxisConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (yAxisConfig *YAxisConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `YAxisConfig` objects.

```go
func (yAxisConfig *YAxisConfig) Equals(other YAxisConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `YAxisConfig` fields for violations and returns them.

```go
func (yAxisConfig *YAxisConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [YAxisConfigBuilder](./builder-YAxisConfigBuilder.md)
