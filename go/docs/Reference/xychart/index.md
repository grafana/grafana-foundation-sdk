# <span class="badge package-variant-panelcfg"></span> xychart

## Objects

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-struct"></span> [ScatterSeriesConfig](./object-ScatterSeriesConfig.md)
 * <span class="badge object-type-enum"></span> [ScatterShow](./object-ScatterShow.md)
 * <span class="badge object-type-enum"></span> [SeriesMapping](./object-SeriesMapping.md)
 * <span class="badge object-type-struct"></span> [XYDimensionConfig](./object-XYDimensionConfig.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> NewXYDimensionConfig

NewXYDimensionConfig creates a new XYDimensionConfig object.

```go
func NewXYDimensionConfig() *XYDimensionConfig
```

### <span class="badge function"></span> NewFieldConfig

NewFieldConfig creates a new FieldConfig object.

```go
func NewFieldConfig() *FieldConfig
```

### <span class="badge function"></span> NewScatterSeriesConfig

NewScatterSeriesConfig creates a new ScatterSeriesConfig object.

```go
func NewScatterSeriesConfig() *ScatterSeriesConfig
```

### <span class="badge function"></span> NewOptions

NewOptions creates a new Options object.

```go
func NewOptions() *Options
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to xychart panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

