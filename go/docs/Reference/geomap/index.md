# geomap

## Objects

 * <span class="badge object-type-struct"></span> [ControlsOptions](./object-ControlsOptions.md)
 * <span class="badge object-type-enum"></span> [MapCenterID](./object-MapCenterID.md)
 * <span class="badge object-type-struct"></span> [MapViewConfig](./object-MapViewConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-enum"></span> [TooltipMode](./object-TooltipMode.md)
 * <span class="badge object-type-struct"></span> [TooltipOptions](./object-TooltipOptions.md)
## Builders

 * <span class="badge builder"></span> [ControlsOptionsBuilder](./builder-ControlsOptionsBuilder.md)
 * <span class="badge builder"></span> [MapViewConfigBuilder](./builder-MapViewConfigBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [TooltipOptionsBuilder](./builder-TooltipOptionsBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to geomap panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> MapViewConfigConverter

MapViewConfigConverter accepts a `MapViewConfig` object and generates the Go code to build this object using builders.

```go
func MapViewConfigConverter(input MapViewConfig) string
```

### <span class="badge function"></span> ControlsOptionsConverter

ControlsOptionsConverter accepts a `ControlsOptions` object and generates the Go code to build this object using builders.

```go
func ControlsOptionsConverter(input ControlsOptions) string
```

### <span class="badge function"></span> TooltipOptionsConverter

TooltipOptionsConverter accepts a `TooltipOptions` object and generates the Go code to build this object using builders.

```go
func TooltipOptionsConverter(input TooltipOptions) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

