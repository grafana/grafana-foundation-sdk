---
title: <span class="badge builder"></span> LineConfigBuilder
---
# <span class="badge builder"></span> LineConfigBuilder

## Constructor

```java
new LineConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public LineConfig build()
```

### <span class="badge object-method"></span> lineColor

```java
public LineConfigBuilder lineColor(String lineColor)
```

### <span class="badge object-method"></span> lineInterpolation

```java
public LineConfigBuilder lineInterpolation(LineInterpolation lineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

```java
public LineConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle)
```

### <span class="badge object-method"></span> lineWidth

```java
public LineConfigBuilder lineWidth(Double lineWidth)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```java
public LineConfigBuilder spanNulls(BoolOrFloat64 spanNulls)
```

## See also

 * <span class="badge object-type-class"></span> [LineConfig](./object-LineConfig.md)
