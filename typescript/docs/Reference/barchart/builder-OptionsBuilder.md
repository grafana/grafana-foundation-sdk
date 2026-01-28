---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```typescript
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> barRadius

Controls the radius of each bar.

```typescript
barRadius(barRadius: number)
```

### <span class="badge object-method"></span> barWidth

Controls the width of bars. 1 = Max width, 0 = Min width.

```typescript
barWidth(barWidth: number)
```

### <span class="badge object-method"></span> colorByField

Use the color value for a sibling field to color each bar value.

```typescript
colorByField(colorByField: string)
```

### <span class="badge object-method"></span> fullHighlight

Enables mode which highlights the entire bar area and shows tooltip when cursor

hovers over highlighted area

```typescript
fullHighlight(fullHighlight: boolean)
```

### <span class="badge object-method"></span> groupWidth

Controls the width of groups. 1 = max with, 0 = min width.

```typescript
groupWidth(groupWidth: number)
```

### <span class="badge object-method"></span> legend

```typescript
legend(legend: cog.Builder<common.VizLegendOptions>)
```

### <span class="badge object-method"></span> orientation

Controls the orientation of the bar chart, either vertical or horizontal.

```typescript
orientation(orientation: common.VizOrientation)
```

### <span class="badge object-method"></span> showValue

This controls whether values are shown on top or to the left of bars.

```typescript
showValue(showValue: common.VisibilityMode)
```

### <span class="badge object-method"></span> stacking

Controls whether bars are stacked or not, either normally or in percent mode.

```typescript
stacking(stacking: common.StackingMode)
```

### <span class="badge object-method"></span> text

```typescript
text(text: cog.Builder<common.VizTextDisplayOptions>)
```

### <span class="badge object-method"></span> tooltip

```typescript
tooltip(tooltip: cog.Builder<common.VizTooltipOptions>)
```

### <span class="badge object-method"></span> xField

Manually select which field from the dataset to represent the x field.

```typescript
xField(xField: string)
```

### <span class="badge object-method"></span> xTickLabelMaxLength

Sets the max length that a label can have before it is truncated.

```typescript
xTickLabelMaxLength(xTickLabelMaxLength: number)
```

### <span class="badge object-method"></span> xTickLabelRotation

Controls the rotation of the x axis labels.

```typescript
xTickLabelRotation(xTickLabelRotation: number)
```

### <span class="badge object-method"></span> xTickLabelSpacing

Controls the spacing between x axis labels.

negative values indicate backwards skipping behavior

```typescript
xTickLabelSpacing(xTickLabelSpacing: number)
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
