# <span class="badge package-variant-panelcfg"></span> canvas

## Objects

 * <span class="badge object-type-struct"></span> [BackgroundConfig](./object-BackgroundConfig.md)
 * <span class="badge object-type-enum"></span> [BackgroundImageSize](./object-BackgroundImageSize.md)
 * <span class="badge object-type-struct"></span> [CanvasConnection](./object-CanvasConnection.md)
 * <span class="badge object-type-struct"></span> [CanvasElementOptions](./object-CanvasElementOptions.md)
 * <span class="badge object-type-struct"></span> [CanvasOptionsRoot](./object-CanvasOptionsRoot.md)
 * <span class="badge object-type-struct"></span> [ConnectionCoordinates](./object-ConnectionCoordinates.md)
 * <span class="badge object-type-enum"></span> [ConnectionPath](./object-ConnectionPath.md)
 * <span class="badge object-type-struct"></span> [Constraint](./object-Constraint.md)
 * <span class="badge object-type-enum"></span> [HorizontalConstraint](./object-HorizontalConstraint.md)
 * <span class="badge object-type-enum"></span> [HttpRequestMethod](./object-HttpRequestMethod.md)
 * <span class="badge object-type-struct"></span> [LineConfig](./object-LineConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-struct"></span> [Placement](./object-Placement.md)
 * <span class="badge object-type-enum"></span> [VerticalConstraint](./object-VerticalConstraint.md)
## Builders

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> NewConstraint

NewConstraint creates a new Constraint object.

```go
func NewConstraint() *Constraint
```

### <span class="badge function"></span> NewPlacement

NewPlacement creates a new Placement object.

```go
func NewPlacement() *Placement
```

### <span class="badge function"></span> NewBackgroundConfig

NewBackgroundConfig creates a new BackgroundConfig object.

```go
func NewBackgroundConfig() *BackgroundConfig
```

### <span class="badge function"></span> NewLineConfig

NewLineConfig creates a new LineConfig object.

```go
func NewLineConfig() *LineConfig
```

### <span class="badge function"></span> NewConnectionCoordinates

NewConnectionCoordinates creates a new ConnectionCoordinates object.

```go
func NewConnectionCoordinates() *ConnectionCoordinates
```

### <span class="badge function"></span> NewCanvasConnection

NewCanvasConnection creates a new CanvasConnection object.

```go
func NewCanvasConnection() *CanvasConnection
```

### <span class="badge function"></span> NewCanvasElementOptions

NewCanvasElementOptions creates a new CanvasElementOptions object.

```go
func NewCanvasElementOptions() *CanvasElementOptions
```

### <span class="badge function"></span> NewOptions

NewOptions creates a new Options object.

```go
func NewOptions() *Options
```

### <span class="badge function"></span> NewCanvasOptionsRoot

NewCanvasOptionsRoot creates a new CanvasOptionsRoot object.

```go
func NewCanvasOptionsRoot() *CanvasOptionsRoot
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to canvas panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

