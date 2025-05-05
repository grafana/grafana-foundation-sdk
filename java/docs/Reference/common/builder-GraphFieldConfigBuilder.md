---
title: <span class="badge builder"></span> GraphFieldConfigBuilder
---
# <span class="badge builder"></span> GraphFieldConfigBuilder

## Constructor

```java
new GraphFieldConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public GraphFieldConfig build()
```

### <span class="badge object-method"></span> axisCenteredZero

```java
public GraphFieldConfigBuilder axisCenteredZero(Boolean axisCenteredZero)
```

### <span class="badge object-method"></span> axisColorMode

```java
public GraphFieldConfigBuilder axisColorMode(AxisColorMode axisColorMode)
```

### <span class="badge object-method"></span> axisGridShow

```java
public GraphFieldConfigBuilder axisGridShow(Boolean axisGridShow)
```

### <span class="badge object-method"></span> axisLabel

```java
public GraphFieldConfigBuilder axisLabel(String axisLabel)
```

### <span class="badge object-method"></span> axisPlacement

```java
public GraphFieldConfigBuilder axisPlacement(AxisPlacement axisPlacement)
```

### <span class="badge object-method"></span> axisSoftMax

```java
public GraphFieldConfigBuilder axisSoftMax(Double axisSoftMax)
```

### <span class="badge object-method"></span> axisSoftMin

```java
public GraphFieldConfigBuilder axisSoftMin(Double axisSoftMin)
```

### <span class="badge object-method"></span> axisWidth

```java
public GraphFieldConfigBuilder axisWidth(Double axisWidth)
```

### <span class="badge object-method"></span> barAlignment

```java
public GraphFieldConfigBuilder barAlignment(BarAlignment barAlignment)
```

### <span class="badge object-method"></span> barMaxWidth

```java
public GraphFieldConfigBuilder barMaxWidth(Double barMaxWidth)
```

### <span class="badge object-method"></span> barWidthFactor

```java
public GraphFieldConfigBuilder barWidthFactor(Double barWidthFactor)
```

### <span class="badge object-method"></span> drawStyle

```java
public GraphFieldConfigBuilder drawStyle(GraphDrawStyle drawStyle)
```

### <span class="badge object-method"></span> fillBelowTo

```java
public GraphFieldConfigBuilder fillBelowTo(String fillBelowTo)
```

### <span class="badge object-method"></span> fillColor

```java
public GraphFieldConfigBuilder fillColor(String fillColor)
```

### <span class="badge object-method"></span> fillOpacity

```java
public GraphFieldConfigBuilder fillOpacity(Double fillOpacity)
```

### <span class="badge object-method"></span> gradientMode

```java
public GraphFieldConfigBuilder gradientMode(GraphGradientMode gradientMode)
```

### <span class="badge object-method"></span> hideFrom

```java
public GraphFieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom)
```

### <span class="badge object-method"></span> insertNulls

```java
public GraphFieldConfigBuilder insertNulls(BoolOrUint32 insertNulls)
```

### <span class="badge object-method"></span> lineColor

```java
public GraphFieldConfigBuilder lineColor(String lineColor)
```

### <span class="badge object-method"></span> lineInterpolation

```java
public GraphFieldConfigBuilder lineInterpolation(LineInterpolation lineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

```java
public GraphFieldConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle)
```

### <span class="badge object-method"></span> lineWidth

```java
public GraphFieldConfigBuilder lineWidth(Double lineWidth)
```

### <span class="badge object-method"></span> pointColor

```java
public GraphFieldConfigBuilder pointColor(String pointColor)
```

### <span class="badge object-method"></span> pointSize

```java
public GraphFieldConfigBuilder pointSize(Double pointSize)
```

### <span class="badge object-method"></span> pointSymbol

```java
public GraphFieldConfigBuilder pointSymbol(String pointSymbol)
```

### <span class="badge object-method"></span> scaleDistribution

```java
public GraphFieldConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution)
```

### <span class="badge object-method"></span> showPoints

```java
public GraphFieldConfigBuilder showPoints(VisibilityMode showPoints)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```java
public GraphFieldConfigBuilder spanNulls(BoolOrFloat64 spanNulls)
```

### <span class="badge object-method"></span> stacking

```java
public GraphFieldConfigBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking)
```

### <span class="badge object-method"></span> thresholdsStyle

```java
public GraphFieldConfigBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle)
```

### <span class="badge object-method"></span> transform

```java
public GraphFieldConfigBuilder transform(GraphTransform transform)
```

## See also

 * <span class="badge object-type-class"></span> [GraphFieldConfig](./object-GraphFieldConfig.md)
