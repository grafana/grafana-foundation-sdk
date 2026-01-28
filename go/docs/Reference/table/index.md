# <span class="badge package-variant-panelcfg"></span> table

## Objects

 * <span class="badge object-type-ref"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
## Builders

 * <span class="badge builder"></span> [FieldConfigBuilder](./builder-FieldConfigBuilder.md)
 * <span class="badge builder"></span> [OptionsBuilder](./builder-OptionsBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
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

VariantConfig returns the configuration related to table panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> OptionsConverter

OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.

```go
func OptionsConverter(input Options) string
```

### <span class="badge function"></span> FieldConfigConverter

FieldConfigConverter accepts a `FieldConfig` object and generates the Go code to build this object using builders.

```go
func FieldConfigConverter(input FieldConfig) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

