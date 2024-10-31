---
title: <span class="badge object-type-struct"></span> FieldConfig
---
# <span class="badge object-type-struct"></span> FieldConfig

## Definition

```go
type FieldConfig struct {
    Show *xychart.ScatterShow `json:"show,omitempty"`
    PointSize *common.ScaleDimensionConfig `json:"pointSize,omitempty"`
    PointColor *common.ColorDimensionConfig `json:"pointColor,omitempty"`
    LineColor *common.ColorDimensionConfig `json:"lineColor,omitempty"`
    LineWidth *int32 `json:"lineWidth,omitempty"`
    LineStyle *common.LineStyle `json:"lineStyle,omitempty"`
    Label *common.VisibilityMode `json:"label,omitempty"`
    HideFrom *common.HideSeriesConfig `json:"hideFrom,omitempty"`
    AxisPlacement *common.AxisPlacement `json:"axisPlacement,omitempty"`
    AxisColorMode *common.AxisColorMode `json:"axisColorMode,omitempty"`
    AxisLabel *string `json:"axisLabel,omitempty"`
    AxisWidth *float64 `json:"axisWidth,omitempty"`
    AxisSoftMin *float64 `json:"axisSoftMin,omitempty"`
    AxisSoftMax *float64 `json:"axisSoftMax,omitempty"`
    AxisGridShow *bool `json:"axisGridShow,omitempty"`
    ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
    AxisCenteredZero *bool `json:"axisCenteredZero,omitempty"`
    LabelValue *common.TextDimensionConfig `json:"labelValue,omitempty"`
    AxisBorderShow *bool `json:"axisBorderShow,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fieldConfig *FieldConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FieldConfig` objects.

```go
func (fieldConfig *FieldConfig) Equals(other FieldConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.

```go
func (fieldConfig *FieldConfig) Validate() error
```

