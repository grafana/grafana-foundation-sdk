---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> BarRadius

Controls the radius of each bar.

```go
func (builder *OptionsBuilder) BarRadius(barRadius float64) *OptionsBuilder
```

### <span class="badge object-method"></span> BarWidth

Controls the width of bars. 1 = Max width, 0 = Min width.

```go
func (builder *OptionsBuilder) BarWidth(barWidth float64) *OptionsBuilder
```

### <span class="badge object-method"></span> ColorByField

Use the color value for a sibling field to color each bar value.

```go
func (builder *OptionsBuilder) ColorByField(colorByField string) *OptionsBuilder
```

### <span class="badge object-method"></span> FullHighlight

Enables mode which highlights the entire bar area and shows tooltip when cursor

hovers over highlighted area

```go
func (builder *OptionsBuilder) FullHighlight(fullHighlight bool) *OptionsBuilder
```

### <span class="badge object-method"></span> GroupWidth

Controls the width of groups. 1 = max with, 0 = min width.

```go
func (builder *OptionsBuilder) GroupWidth(groupWidth float64) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Orientation

Controls the orientation of the bar chart, either vertical or horizontal.

```go
func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowValue

This controls whether values are shown on top or to the left of bars.

```go
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder
```

### <span class="badge object-method"></span> Stacking

Controls whether bars are stacked or not, either normally or in percent mode.

```go
func (builder *OptionsBuilder) Stacking(stacking common.StackingMode) *OptionsBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> XField

Manually select which field from the dataset to represent the x field.

```go
func (builder *OptionsBuilder) XField(xField string) *OptionsBuilder
```

### <span class="badge object-method"></span> XTickLabelMaxLength

Sets the max length that a label can have before it is truncated.

```go
func (builder *OptionsBuilder) XTickLabelMaxLength(xTickLabelMaxLength int32) *OptionsBuilder
```

### <span class="badge object-method"></span> XTickLabelRotation

Controls the rotation of the x axis labels.

```go
func (builder *OptionsBuilder) XTickLabelRotation(xTickLabelRotation int32) *OptionsBuilder
```

### <span class="badge object-method"></span> XTickLabelSpacing

Controls the spacing between x axis labels.

negative values indicate backwards skipping behavior

```go
func (builder *OptionsBuilder) XTickLabelSpacing(xTickLabelSpacing int32) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
