# <span class="badge package-variant-panelcfg"></span> heatmap

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

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> NewHeatmapColorOptions

NewHeatmapColorOptions creates a new HeatmapColorOptions object.

```go
func NewHeatmapColorOptions() *HeatmapColorOptions
```

### <span class="badge function"></span> NewYAxisConfig

NewYAxisConfig creates a new YAxisConfig object.

```go
func NewYAxisConfig() *YAxisConfig
```

### <span class="badge function"></span> NewCellValues

NewCellValues creates a new CellValues object.

```go
func NewCellValues() *CellValues
```

### <span class="badge function"></span> NewFilterValueRange

NewFilterValueRange creates a new FilterValueRange object.

```go
func NewFilterValueRange() *FilterValueRange
```

### <span class="badge function"></span> NewHeatmapTooltip

NewHeatmapTooltip creates a new HeatmapTooltip object.

```go
func NewHeatmapTooltip() *HeatmapTooltip
```

### <span class="badge function"></span> NewHeatmapLegend

NewHeatmapLegend creates a new HeatmapLegend object.

```go
func NewHeatmapLegend() *HeatmapLegend
```

### <span class="badge function"></span> NewExemplarConfig

NewExemplarConfig creates a new ExemplarConfig object.

```go
func NewExemplarConfig() *ExemplarConfig
```

### <span class="badge function"></span> NewRowsHeatmapOptions

NewRowsHeatmapOptions creates a new RowsHeatmapOptions object.

```go
func NewRowsHeatmapOptions() *RowsHeatmapOptions
```

### <span class="badge function"></span> NewOptions

NewOptions creates a new Options object.

```go
func NewOptions() *Options
```

### <span class="badge function"></span> NewFieldConfig

NewFieldConfig creates a new FieldConfig object.

```go
func NewFieldConfig() *FieldConfig
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to heatmap panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

