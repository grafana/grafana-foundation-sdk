# text

## Objects

 * <span class="badge object-type-enum"></span> [CodeLanguage](./object-CodeLanguage.md)
 * <span class="badge object-type-struct"></span> [CodeOptions](./object-CodeOptions.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-enum"></span> [TextMode](./object-TextMode.md)
## Builders

 * <span class="badge builder"></span> [CodeOptionsBuilder](./builder-CodeOptionsBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to text panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> CodeOptionsConverter

CodeOptionsConverter accepts a `CodeOptions` object and generates the Go code to build this object using builders.

```go
func CodeOptionsConverter(input CodeOptions) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

