# piechart

## Objects

 * <span class="badge object-type-ref"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-enum"></span> [PieChartLabels](./object-PieChartLabels.md)
 * <span class="badge object-type-struct"></span> [PieChartLegendOptions](./object-PieChartLegendOptions.md)
 * <span class="badge object-type-enum"></span> [PieChartLegendValues](./object-PieChartLegendValues.md)
 * <span class="badge object-type-enum"></span> [PieChartType](./object-PieChartType.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [PieChartLegendOptionsBuilder](./builder-PieChartLegendOptionsBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to piechart panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> PieChartLegendOptionsConverter

PieChartLegendOptionsConverter accepts a `PieChartLegendOptions` object and generates the Go code to build this object using builders.

```go
func PieChartLegendOptionsConverter(input PieChartLegendOptions) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

