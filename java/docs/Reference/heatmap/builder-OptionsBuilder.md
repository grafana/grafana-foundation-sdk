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

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```java
public OptionsBuilder calculate(Boolean calculate)
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

```java
public OptionsBuilder calculation(com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> calculation)
```

### <span class="badge object-method"></span> cellGap

Controls gap between cells

```java
public OptionsBuilder cellGap(Integer cellGap)
```

### <span class="badge object-method"></span> cellRadius

Controls cell radius

```java
public OptionsBuilder cellRadius(Float cellRadius)
```

### <span class="badge object-method"></span> cellValues

Controls cell value unit

```java
public OptionsBuilder cellValues(com.grafana.foundation.cog.Builder<CellValues> cellValues)
```

### <span class="badge object-method"></span> color

Controls the color options

```java
public OptionsBuilder color(com.grafana.foundation.cog.Builder<HeatmapColorOptions> color)
```

### <span class="badge object-method"></span> exemplars

Controls exemplar options

```java
public OptionsBuilder exemplars(com.grafana.foundation.cog.Builder<ExemplarConfig> exemplars)
```

### <span class="badge object-method"></span> filterValues

Filters values between a given range

```java
public OptionsBuilder filterValues(com.grafana.foundation.cog.Builder<FilterValueRange> filterValues)
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```java
public OptionsBuilder legend(com.grafana.foundation.cog.Builder<HeatmapLegend> legend)
```

### <span class="badge object-method"></span> rowsFrame

Controls tick alignment and value name when not calculating from data

```java
public OptionsBuilder rowsFrame(com.grafana.foundation.cog.Builder<RowsHeatmapOptions> rowsFrame)
```

### <span class="badge object-method"></span> showValue

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```java
public OptionsBuilder showValue(VisibilityMode showValue)
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

```java
public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<HeatmapTooltip> tooltip)
```

### <span class="badge object-method"></span> yAxis

Controls yAxis placement

```java
public OptionsBuilder yAxis(com.grafana.foundation.cog.Builder<YAxisConfig> yAxis)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
