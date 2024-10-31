# nodegraph

## Objects

 * <span class="badge object-type-struct"></span> [ArcOption](./object-ArcOption.md)
 * <span class="badge object-type-struct"></span> [EdgeOptions](./object-EdgeOptions.md)
 * <span class="badge object-type-struct"></span> [NodeOptions](./object-NodeOptions.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
## Builders

 * <span class="badge builder"></span> [ArcOptionBuilder](./builder-ArcOptionBuilder.md)
 * <span class="badge builder"></span> [EdgeOptionsBuilder](./builder-EdgeOptionsBuilder.md)
 * <span class="badge builder"></span> [NodeOptionsBuilder](./builder-NodeOptionsBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to nodegraph panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> ArcOptionConverter

ArcOptionConverter accepts a `ArcOption` object and generates the Go code to build this object using builders.

```go
func ArcOptionConverter(input ArcOption) string
```

### <span class="badge function"></span> NodeOptionsConverter

NodeOptionsConverter accepts a `NodeOptions` object and generates the Go code to build this object using builders.

```go
func NodeOptionsConverter(input NodeOptions) string
```

### <span class="badge function"></span> EdgeOptionsConverter

EdgeOptionsConverter accepts a `EdgeOptions` object and generates the Go code to build this object using builders.

```go
func EdgeOptionsConverter(input EdgeOptions) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

