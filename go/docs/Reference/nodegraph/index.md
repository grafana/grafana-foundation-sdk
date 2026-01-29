# <span class="badge package-variant-panelcfg"></span> nodegraph

## Objects

 * <span class="badge object-type-struct"></span> [ArcOption](./object-ArcOption.md)
 * <span class="badge object-type-struct"></span> [EdgeOptions](./object-EdgeOptions.md)
 * <span class="badge object-type-struct"></span> [NodeOptions](./object-NodeOptions.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-enum"></span> [ZoomMode](./object-ZoomMode.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> NewArcOption

NewArcOption creates a new ArcOption object.

```go
func NewArcOption() *ArcOption
```

### <span class="badge function"></span> NewNodeOptions

NewNodeOptions creates a new NodeOptions object.

```go
func NewNodeOptions() *NodeOptions
```

### <span class="badge function"></span> NewEdgeOptions

NewEdgeOptions creates a new EdgeOptions object.

```go
func NewEdgeOptions() *EdgeOptions
```

### <span class="badge function"></span> NewOptions

NewOptions creates a new Options object.

```go
func NewOptions() *Options
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to nodegraph panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

