# xychart

## Objects

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [MatcherConfig](./object-MatcherConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-enum"></span> [PointShape](./object-PointShape.md)
 * <span class="badge object-type-enum"></span> [SeriesMapping](./object-SeriesMapping.md)
 * <span class="badge object-type-struct"></span> [XYSeriesConfig](./object-XYSeriesConfig.md)
 * <span class="badge object-type-enum"></span> [XYShowMode](./object-XYShowMode.md)
 * <span class="badge object-type-struct"></span> [XychartFieldConfigPointSize](./object-XychartFieldConfigPointSize.md)
 * <span class="badge object-type-struct"></span> [XychartXYSeriesConfigColor](./object-XychartXYSeriesConfigColor.md)
 * <span class="badge object-type-struct"></span> [XychartXYSeriesConfigFrame](./object-XychartXYSeriesConfigFrame.md)
 * <span class="badge object-type-struct"></span> [XychartXYSeriesConfigName](./object-XychartXYSeriesConfigName.md)
 * <span class="badge object-type-struct"></span> [XychartXYSeriesConfigSize](./object-XychartXYSeriesConfigSize.md)
 * <span class="badge object-type-struct"></span> [XychartXYSeriesConfigX](./object-XychartXYSeriesConfigX.md)
 * <span class="badge object-type-struct"></span> [XychartXYSeriesConfigY](./object-XychartXYSeriesConfigY.md)
## Builders

 * <span class="badge builder"></span> [MatcherConfigBuilder](./builder-MatcherConfigBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [XYSeriesConfigBuilder](./builder-XYSeriesConfigBuilder.md)
 * <span class="badge builder"></span> [XychartFieldConfigPointSizeBuilder](./builder-XychartFieldConfigPointSizeBuilder.md)
 * <span class="badge builder"></span> [XychartXYSeriesConfigColorBuilder](./builder-XychartXYSeriesConfigColorBuilder.md)
 * <span class="badge builder"></span> [XychartXYSeriesConfigFrameBuilder](./builder-XychartXYSeriesConfigFrameBuilder.md)
 * <span class="badge builder"></span> [XychartXYSeriesConfigNameBuilder](./builder-XychartXYSeriesConfigNameBuilder.md)
 * <span class="badge builder"></span> [XychartXYSeriesConfigSizeBuilder](./builder-XychartXYSeriesConfigSizeBuilder.md)
 * <span class="badge builder"></span> [XychartXYSeriesConfigXBuilder](./builder-XychartXYSeriesConfigXBuilder.md)
 * <span class="badge builder"></span> [XychartXYSeriesConfigYBuilder](./builder-XychartXYSeriesConfigYBuilder.md)
## Functions

### <span class="badge function"></span> NewMatcherConfig

NewMatcherConfig creates a new MatcherConfig object.

```go
func NewMatcherConfig() *MatcherConfig
```

### <span class="badge function"></span> NewFieldConfig

NewFieldConfig creates a new FieldConfig object.

```go
func NewFieldConfig() *FieldConfig
```

### <span class="badge function"></span> NewXYSeriesConfig

NewXYSeriesConfig creates a new XYSeriesConfig object.

```go
func NewXYSeriesConfig() *XYSeriesConfig
```

### <span class="badge function"></span> NewOptions

NewOptions creates a new Options object.

```go
func NewOptions() *Options
```

### <span class="badge function"></span> NewXychartFieldConfigPointSize

NewXychartFieldConfigPointSize creates a new XychartFieldConfigPointSize object.

```go
func NewXychartFieldConfigPointSize() *XychartFieldConfigPointSize
```

### <span class="badge function"></span> NewXychartXYSeriesConfigName

NewXychartXYSeriesConfigName creates a new XychartXYSeriesConfigName object.

```go
func NewXychartXYSeriesConfigName() *XychartXYSeriesConfigName
```

### <span class="badge function"></span> NewXychartXYSeriesConfigFrame

NewXychartXYSeriesConfigFrame creates a new XychartXYSeriesConfigFrame object.

```go
func NewXychartXYSeriesConfigFrame() *XychartXYSeriesConfigFrame
```

### <span class="badge function"></span> NewXychartXYSeriesConfigX

NewXychartXYSeriesConfigX creates a new XychartXYSeriesConfigX object.

```go
func NewXychartXYSeriesConfigX() *XychartXYSeriesConfigX
```

### <span class="badge function"></span> NewXychartXYSeriesConfigY

NewXychartXYSeriesConfigY creates a new XychartXYSeriesConfigY object.

```go
func NewXychartXYSeriesConfigY() *XychartXYSeriesConfigY
```

### <span class="badge function"></span> NewXychartXYSeriesConfigColor

NewXychartXYSeriesConfigColor creates a new XychartXYSeriesConfigColor object.

```go
func NewXychartXYSeriesConfigColor() *XychartXYSeriesConfigColor
```

### <span class="badge function"></span> NewXychartXYSeriesConfigSize

NewXychartXYSeriesConfigSize creates a new XychartXYSeriesConfigSize object.

```go
func NewXychartXYSeriesConfigSize() *XychartXYSeriesConfigSize
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to xychart panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> MatcherConfigConverter

MatcherConfigConverter accepts a `MatcherConfig` object and generates the Go code to build this object using builders.

```go
func MatcherConfigConverter(input MatcherConfig) string
```

### <span class="badge function"></span> XYSeriesConfigConverter

XYSeriesConfigConverter accepts a `XYSeriesConfig` object and generates the Go code to build this object using builders.

```go
func XYSeriesConfigConverter(input XYSeriesConfig) string
```

### <span class="badge function"></span> XychartFieldConfigPointSizeConverter

XychartFieldConfigPointSizeConverter accepts a `XychartFieldConfigPointSize` object and generates the Go code to build this object using builders.

```go
func XychartFieldConfigPointSizeConverter(input XychartFieldConfigPointSize) string
```

### <span class="badge function"></span> XychartXYSeriesConfigNameConverter

XychartXYSeriesConfigNameConverter accepts a `XychartXYSeriesConfigName` object and generates the Go code to build this object using builders.

```go
func XychartXYSeriesConfigNameConverter(input XychartXYSeriesConfigName) string
```

### <span class="badge function"></span> XychartXYSeriesConfigFrameConverter

XychartXYSeriesConfigFrameConverter accepts a `XychartXYSeriesConfigFrame` object and generates the Go code to build this object using builders.

```go
func XychartXYSeriesConfigFrameConverter(input XychartXYSeriesConfigFrame) string
```

### <span class="badge function"></span> XychartXYSeriesConfigXConverter

XychartXYSeriesConfigXConverter accepts a `XychartXYSeriesConfigX` object and generates the Go code to build this object using builders.

```go
func XychartXYSeriesConfigXConverter(input XychartXYSeriesConfigX) string
```

### <span class="badge function"></span> XychartXYSeriesConfigYConverter

XychartXYSeriesConfigYConverter accepts a `XychartXYSeriesConfigY` object and generates the Go code to build this object using builders.

```go
func XychartXYSeriesConfigYConverter(input XychartXYSeriesConfigY) string
```

### <span class="badge function"></span> XychartXYSeriesConfigColorConverter

XychartXYSeriesConfigColorConverter accepts a `XychartXYSeriesConfigColor` object and generates the Go code to build this object using builders.

```go
func XychartXYSeriesConfigColorConverter(input XychartXYSeriesConfigColor) string
```

### <span class="badge function"></span> XychartXYSeriesConfigSizeConverter

XychartXYSeriesConfigSizeConverter accepts a `XychartXYSeriesConfigSize` object and generates the Go code to build this object using builders.

```go
func XychartXYSeriesConfigSizeConverter(input XychartXYSeriesConfigSize) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

