---
title: <span class="badge builder"></span> FieldConfigBuilder
---
# <span class="badge builder"></span> FieldConfigBuilder

## Constructor

```java
new FieldConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public FieldConfig build()
```

### <span class="badge object-method"></span> axisCenteredZero

```java
public FieldConfigBuilder axisCenteredZero(Boolean axisCenteredZero)
```

### <span class="badge object-method"></span> axisColorMode

```java
public FieldConfigBuilder axisColorMode(AxisColorMode axisColorMode)
```

### <span class="badge object-method"></span> axisGridShow

```java
public FieldConfigBuilder axisGridShow(Boolean axisGridShow)
```

### <span class="badge object-method"></span> axisLabel

```java
public FieldConfigBuilder axisLabel(String axisLabel)
```

### <span class="badge object-method"></span> axisPlacement

```java
public FieldConfigBuilder axisPlacement(AxisPlacement axisPlacement)
```

### <span class="badge object-method"></span> axisSoftMax

```java
public FieldConfigBuilder axisSoftMax(Double axisSoftMax)
```

### <span class="badge object-method"></span> axisSoftMin

```java
public FieldConfigBuilder axisSoftMin(Double axisSoftMin)
```

### <span class="badge object-method"></span> axisWidth

```java
public FieldConfigBuilder axisWidth(Double axisWidth)
```

### <span class="badge object-method"></span> barAlignment

```java
public FieldConfigBuilder barAlignment(BarAlignment barAlignment)
```

### <span class="badge object-method"></span> barMaxWidth

```java
public FieldConfigBuilder barMaxWidth(Double barMaxWidth)
```

### <span class="badge object-method"></span> barWidthFactor

```java
public FieldConfigBuilder barWidthFactor(Double barWidthFactor)
```

### <span class="badge object-method"></span> drawStyle

```java
public FieldConfigBuilder drawStyle(GraphDrawStyle drawStyle)
```

### <span class="badge object-method"></span> fillBelowTo

```java
public FieldConfigBuilder fillBelowTo(String fillBelowTo)
```

### <span class="badge object-method"></span> fillColor

```java
public FieldConfigBuilder fillColor(String fillColor)
```

### <span class="badge object-method"></span> fillOpacity

```java
public FieldConfigBuilder fillOpacity(Double fillOpacity)
```

### <span class="badge object-method"></span> gradientMode

```java
public FieldConfigBuilder gradientMode(GraphGradientMode gradientMode)
```

### <span class="badge object-method"></span> hideFrom

```java
public FieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom)
```

### <span class="badge object-method"></span> insertNulls

```java
public FieldConfigBuilder insertNulls(BoolOrUint32 insertNulls)
```

### <span class="badge object-method"></span> lineColor

```java
public FieldConfigBuilder lineColor(String lineColor)
```

### <span class="badge object-method"></span> lineInterpolation

```java
public FieldConfigBuilder lineInterpolation(LineInterpolation lineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

```java
public FieldConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle)
```

### <span class="badge object-method"></span> lineWidth

```java
public FieldConfigBuilder lineWidth(Double lineWidth)
```

### <span class="badge object-method"></span> pointColor

```java
public FieldConfigBuilder pointColor(String pointColor)
```

### <span class="badge object-method"></span> pointSize

```java
public FieldConfigBuilder pointSize(Double pointSize)
```

### <span class="badge object-method"></span> pointSymbol

```java
public FieldConfigBuilder pointSymbol(String pointSymbol)
```

### <span class="badge object-method"></span> scaleDistribution

```java
public FieldConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution)
```

### <span class="badge object-method"></span> showPoints

```java
public FieldConfigBuilder showPoints(VisibilityMode showPoints)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```java
public FieldConfigBuilder spanNulls(BoolOrFloat64 spanNulls)
```

### <span class="badge object-method"></span> stacking

```java
public FieldConfigBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking)
```

### <span class="badge object-method"></span> thresholdsStyle

```java
public FieldConfigBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle)
```

### <span class="badge object-method"></span> transform

```java
public FieldConfigBuilder transform(GraphTransform transform)
```

## See also

 * <span class="badge object-type-ref"></span> [FieldConfig](./object-FieldConfig.md)
