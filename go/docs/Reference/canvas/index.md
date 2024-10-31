# canvas

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

 * <span class="badge builder"></span> [BackgroundConfigBuilder](./builder-BackgroundConfigBuilder.md)
 * <span class="badge builder"></span> [CanvasConnectionBuilder](./builder-CanvasConnectionBuilder.md)
 * <span class="badge builder"></span> [CanvasElementOptionsBuilder](./builder-CanvasElementOptionsBuilder.md)
 * <span class="badge builder"></span> [CanvasOptionsRootBuilder](./builder-CanvasOptionsRootBuilder.md)
 * <span class="badge builder"></span> [ConnectionCoordinatesBuilder](./builder-ConnectionCoordinatesBuilder.md)
 * <span class="badge builder"></span> [ConstraintBuilder](./builder-ConstraintBuilder.md)
 * <span class="badge builder"></span> [LineConfigBuilder](./builder-LineConfigBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [PlacementBuilder](./builder-PlacementBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to canvas panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> ConstraintConverter

ConstraintConverter accepts a `Constraint` object and generates the Go code to build this object using builders.

```go
func ConstraintConverter(input Constraint) string
```

### <span class="badge function"></span> PlacementConverter

PlacementConverter accepts a `Placement` object and generates the Go code to build this object using builders.

```go
func PlacementConverter(input Placement) string
```

### <span class="badge function"></span> BackgroundConfigConverter

BackgroundConfigConverter accepts a `BackgroundConfig` object and generates the Go code to build this object using builders.

```go
func BackgroundConfigConverter(input BackgroundConfig) string
```

### <span class="badge function"></span> LineConfigConverter

LineConfigConverter accepts a `LineConfig` object and generates the Go code to build this object using builders.

```go
func LineConfigConverter(input LineConfig) string
```

### <span class="badge function"></span> ConnectionCoordinatesConverter

ConnectionCoordinatesConverter accepts a `ConnectionCoordinates` object and generates the Go code to build this object using builders.

```go
func ConnectionCoordinatesConverter(input ConnectionCoordinates) string
```

### <span class="badge function"></span> CanvasConnectionConverter

CanvasConnectionConverter accepts a `CanvasConnection` object and generates the Go code to build this object using builders.

```go
func CanvasConnectionConverter(input CanvasConnection) string
```

### <span class="badge function"></span> CanvasElementOptionsConverter

CanvasElementOptionsConverter accepts a `CanvasElementOptions` object and generates the Go code to build this object using builders.

```go
func CanvasElementOptionsConverter(input CanvasElementOptions) string
```

### <span class="badge function"></span> CanvasOptionsRootConverter

CanvasOptionsRootConverter accepts a `CanvasOptionsRoot` object and generates the Go code to build this object using builders.

```go
func CanvasOptionsRootConverter(input CanvasOptionsRoot) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

