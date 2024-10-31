# xychart

## Objects

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-struct"></span> [ScatterSeriesConfig](./object-ScatterSeriesConfig.md)
 * <span class="badge object-type-enum"></span> [ScatterShow](./object-ScatterShow.md)
 * <span class="badge object-type-enum"></span> [SeriesMapping](./object-SeriesMapping.md)
 * <span class="badge object-type-struct"></span> [XYDimensionConfig](./object-XYDimensionConfig.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [ScatterSeriesConfigBuilder](./builder-ScatterSeriesConfigBuilder.md)
 * <span class="badge builder"></span> [XYDimensionConfigBuilder](./builder-XYDimensionConfigBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to xychart panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> XYDimensionConfigConverter

XYDimensionConfigConverter accepts a `XYDimensionConfig` object and generates the Go code to build this object using builders.

```go
func XYDimensionConfigConverter(input XYDimensionConfig) string
```

### <span class="badge function"></span> ScatterSeriesConfigConverter

ScatterSeriesConfigConverter accepts a `ScatterSeriesConfig` object and generates the Go code to build this object using builders.

```go
func ScatterSeriesConfigConverter(input ScatterSeriesConfig) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

