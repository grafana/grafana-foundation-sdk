# <span class="badge package-variant-panelcfg"></span> histogram

## Objects

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [VisualizationBuilder](./builder-VisualizationBuilder.md)
## Functions

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

VariantConfig returns the configuration related to histogram panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

### <span class="badge function"></span> VisualizationConverter

VisualizationConverter accepts a `Visualization` object and generates the Go code to build this object using builders.

```go
func VisualizationConverter(input dashboardv2beta1.VizConfigKind) string
```

