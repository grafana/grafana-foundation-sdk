# heatmap

## Objects

 * <span class="badge object-type-struct"></span> [CellValues](./object-CellValues.md)
 * <span class="badge object-type-struct"></span> [ExemplarConfig](./object-ExemplarConfig.md)
 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [FilterValueRange](./object-FilterValueRange.md)
 * <span class="badge object-type-enum"></span> [HeatmapColorMode](./object-HeatmapColorMode.md)
 * <span class="badge object-type-struct"></span> [HeatmapColorOptions](./object-HeatmapColorOptions.md)
 * <span class="badge object-type-enum"></span> [HeatmapColorScale](./object-HeatmapColorScale.md)
 * <span class="badge object-type-struct"></span> [HeatmapLegend](./object-HeatmapLegend.md)
 * <span class="badge object-type-struct"></span> [HeatmapTooltip](./object-HeatmapTooltip.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-struct"></span> [RowsHeatmapOptions](./object-RowsHeatmapOptions.md)
 * <span class="badge object-type-struct"></span> [YAxisConfig](./object-YAxisConfig.md)
## Builders

 * <span class="badge builder"></span> [CellValuesBuilder](./builder-CellValuesBuilder.md)
 * <span class="badge builder"></span> [ExemplarConfigBuilder](./builder-ExemplarConfigBuilder.md)
 * <span class="badge builder"></span> [FilterValueRangeBuilder](./builder-FilterValueRangeBuilder.md)
 * <span class="badge builder"></span> [HeatmapColorOptionsBuilder](./builder-HeatmapColorOptionsBuilder.md)
 * <span class="badge builder"></span> [HeatmapLegendBuilder](./builder-HeatmapLegendBuilder.md)
 * <span class="badge builder"></span> [HeatmapTooltipBuilder](./builder-HeatmapTooltipBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [RowsHeatmapOptionsBuilder](./builder-RowsHeatmapOptionsBuilder.md)
 * <span class="badge builder"></span> [YAxisConfigBuilder](./builder-YAxisConfigBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to heatmap panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> HeatmapColorOptionsConverter

HeatmapColorOptionsConverter accepts a `HeatmapColorOptions` object and generates the Go code to build this object using builders.

```go
func HeatmapColorOptionsConverter(input HeatmapColorOptions) string
```

### <span class="badge function"></span> YAxisConfigConverter

YAxisConfigConverter accepts a `YAxisConfig` object and generates the Go code to build this object using builders.

```go
func YAxisConfigConverter(input YAxisConfig) string
```

### <span class="badge function"></span> CellValuesConverter

CellValuesConverter accepts a `CellValues` object and generates the Go code to build this object using builders.

```go
func CellValuesConverter(input CellValues) string
```

### <span class="badge function"></span> FilterValueRangeConverter

FilterValueRangeConverter accepts a `FilterValueRange` object and generates the Go code to build this object using builders.

```go
func FilterValueRangeConverter(input FilterValueRange) string
```

### <span class="badge function"></span> HeatmapTooltipConverter

HeatmapTooltipConverter accepts a `HeatmapTooltip` object and generates the Go code to build this object using builders.

```go
func HeatmapTooltipConverter(input HeatmapTooltip) string
```

### <span class="badge function"></span> HeatmapLegendConverter

HeatmapLegendConverter accepts a `HeatmapLegend` object and generates the Go code to build this object using builders.

```go
func HeatmapLegendConverter(input HeatmapLegend) string
```

### <span class="badge function"></span> ExemplarConfigConverter

ExemplarConfigConverter accepts a `ExemplarConfig` object and generates the Go code to build this object using builders.

```go
func ExemplarConfigConverter(input ExemplarConfig) string
```

### <span class="badge function"></span> RowsHeatmapOptionsConverter

RowsHeatmapOptionsConverter accepts a `RowsHeatmapOptions` object and generates the Go code to build this object using builders.

```go
func RowsHeatmapOptionsConverter(input RowsHeatmapOptions) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

