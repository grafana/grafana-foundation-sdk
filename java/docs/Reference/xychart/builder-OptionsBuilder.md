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

### <span class="badge object-method"></span> dims

Table Mode (auto)

```java
public OptionsBuilder dims(com.grafana.foundation.cog.Builder<XYDimensionConfig> dims)
```

### <span class="badge object-method"></span> legend

```java
public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> series

Manual Mode

```java
public OptionsBuilder series(List<com.grafana.foundation.cog.Builder<ScatterSeriesConfig>> series)
```

### <span class="badge object-method"></span> seriesMapping

```java
public OptionsBuilder seriesMapping(SeriesMapping seriesMapping)
```

### <span class="badge object-method"></span> tooltip

```java
public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
