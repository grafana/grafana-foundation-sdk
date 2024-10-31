# debug

## Objects

 * <span class="badge object-type-enum"></span> [DebugMode](./object-DebugMode.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-struct"></span> [UpdateConfig](./object-UpdateConfig.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [UpdateConfigBuilder](./builder-UpdateConfigBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to debug panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> UpdateConfigConverter

UpdateConfigConverter accepts a `UpdateConfig` object and generates the Go code to build this object using builders.

```go
func UpdateConfigConverter(input UpdateConfig) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

