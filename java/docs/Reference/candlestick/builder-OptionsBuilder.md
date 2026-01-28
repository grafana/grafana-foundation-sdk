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

### <span class="badge object-method"></span> candleStyle

Sets the style of the candlesticks

```java
public OptionsBuilder candleStyle(CandleStyle candleStyle)
```

### <span class="badge object-method"></span> colorStrategy

Sets the color strategy for the candlesticks

```java
public OptionsBuilder colorStrategy(ColorStrategy colorStrategy)
```

### <span class="badge object-method"></span> colors

Set which colors are used when the price movement is up or down

```java
public OptionsBuilder colors(com.grafana.foundation.cog.Builder<CandlestickColors> colors)
```

### <span class="badge object-method"></span> fields

Map fields to appropriate dimension

```java
public OptionsBuilder fields(com.grafana.foundation.cog.Builder<CandlestickFieldMap> fields)
```

### <span class="badge object-method"></span> includeAllFields

When enabled, all fields will be sent to the graph

```java
public OptionsBuilder includeAllFields(Boolean includeAllFields)
```

### <span class="badge object-method"></span> legend

```java
public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> mode

Sets which dimensions are used for the visualization

```java
public OptionsBuilder mode(VizDisplayMode mode)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
