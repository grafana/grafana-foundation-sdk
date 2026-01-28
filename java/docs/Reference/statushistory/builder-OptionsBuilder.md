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

### <span class="badge object-method"></span> colWidth

Controls the column width

```java
public OptionsBuilder colWidth(Double colWidth)
```

### <span class="badge object-method"></span> legend

```java
public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> rowHeight

Set the height of the rows

```java
public OptionsBuilder rowHeight(Float rowHeight)
```

### <span class="badge object-method"></span> showValue

Show values on the columns

```java
public OptionsBuilder showValue(VisibilityMode showValue)
```

### <span class="badge object-method"></span> timezone

```java
public OptionsBuilder timezone(List<String> timezone)
```

### <span class="badge object-method"></span> tooltip

```java
public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
