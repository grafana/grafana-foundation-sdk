---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```java
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Options build()
```

### <span class="badge object-method"></span> barRadius

Controls the radius of each bar.

```java
public OptionsBuilder barRadius(Double barRadius)
```

### <span class="badge object-method"></span> barWidth

Controls the width of bars. 1 = Max width, 0 = Min width.

```java
public OptionsBuilder barWidth(Double barWidth)
```

### <span class="badge object-method"></span> colorByField

Use the color value for a sibling field to color each bar value.

```java
public OptionsBuilder colorByField(String colorByField)
```

### <span class="badge object-method"></span> fullHighlight

Enables mode which highlights the entire bar area and shows tooltip when cursor

hovers over highlighted area

```java
public OptionsBuilder fullHighlight(Boolean fullHighlight)
```

### <span class="badge object-method"></span> groupWidth

Controls the width of groups. 1 = max with, 0 = min width.

```java
public OptionsBuilder groupWidth(Double groupWidth)
```

### <span class="badge object-method"></span> legend

```java
public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> orientation

Controls the orientation of the bar chart, either vertical or horizontal.

```java
public OptionsBuilder orientation(VizOrientation orientation)
```

### <span class="badge object-method"></span> showValue

This controls whether values are shown on top or to the left of bars.

```java
public OptionsBuilder showValue(VisibilityMode showValue)
```

### <span class="badge object-method"></span> stacking

Controls whether bars are stacked or not, either normally or in percent mode.

```java
public OptionsBuilder stacking(StackingMode stacking)
```

### <span class="badge object-method"></span> text

```java
public OptionsBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text)
```

### <span class="badge object-method"></span> tooltip

```java
public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip)
```

### <span class="badge object-method"></span> xField

Manually select which field from the dataset to represent the x field.

```java
public OptionsBuilder xField(String xField)
```

### <span class="badge object-method"></span> xTickLabelMaxLength

Sets the max length that a label can have before it is truncated.

```java
public OptionsBuilder xTickLabelMaxLength(Integer xTickLabelMaxLength)
```

### <span class="badge object-method"></span> xTickLabelRotation

Controls the rotation of the x axis labels.

```java
public OptionsBuilder xTickLabelRotation(Integer xTickLabelRotation)
```

### <span class="badge object-method"></span> xTickLabelSpacing

Controls the spacing between x axis labels.

negative values indicate backwards skipping behavior

```java
public OptionsBuilder xTickLabelSpacing(Integer xTickLabelSpacing)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
