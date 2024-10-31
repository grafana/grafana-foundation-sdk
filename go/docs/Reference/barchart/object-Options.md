---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    // Manually select which field from the dataset to represent the x field.
    XField *string `json:"xField,omitempty"`
    // Use the color value for a sibling field to color each bar value.
    ColorByField *string `json:"colorByField,omitempty"`
    // Controls the orientation of the bar chart, either vertical or horizontal.
    Orientation common.VizOrientation `json:"orientation"`
    // Controls the radius of each bar.
    BarRadius *float64 `json:"barRadius,omitempty"`
    // Controls the rotation of the x axis labels.
    XTickLabelRotation int32 `json:"xTickLabelRotation"`
    // Sets the max length that a label can have before it is truncated.
    XTickLabelMaxLength int32 `json:"xTickLabelMaxLength"`
    // Controls the spacing between x axis labels.
    // negative values indicate backwards skipping behavior
    XTickLabelSpacing *int32 `json:"xTickLabelSpacing,omitempty"`
    // Controls whether bars are stacked or not, either normally or in percent mode.
    Stacking common.StackingMode `json:"stacking"`
    // This controls whether values are shown on top or to the left of bars.
    ShowValue common.VisibilityMode `json:"showValue"`
    // Controls the width of bars. 1 = Max width, 0 = Min width.
    BarWidth float64 `json:"barWidth"`
    // Controls the width of groups. 1 = max with, 0 = min width.
    GroupWidth float64 `json:"groupWidth"`
    Legend common.VizLegendOptions `json:"legend"`
    Tooltip common.VizTooltipOptions `json:"tooltip"`
    Text *common.VizTextDisplayOptions `json:"text,omitempty"`
    // Enables mode which highlights the entire bar area and shows tooltip when cursor
    // hovers over highlighted area
    FullHighlight bool `json:"fullHighlight"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (options *Options) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Options` objects.

```go
func (options *Options) Equals(other Options) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.

```go
func (options *Options) Validate() error
```

