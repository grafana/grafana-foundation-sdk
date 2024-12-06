# candlestick

## Objects

 * <span class="badge object-type-enum"></span> [CandleStyle](./object-CandleStyle.md)
 * <span class="badge object-type-struct"></span> [CandlestickColors](./object-CandlestickColors.md)
 * <span class="badge object-type-struct"></span> [CandlestickFieldMap](./object-CandlestickFieldMap.md)
 * <span class="badge object-type-enum"></span> [ColorStrategy](./object-ColorStrategy.md)
 * <span class="badge object-type-ref"></span> [FieldConfig](./object-FieldConfig.md)
 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
 * <span class="badge object-type-enum"></span> [VizDisplayMode](./object-VizDisplayMode.md)
## Builders

 * <span class="badge builder"></span> [CandlestickColorsBuilder](./builder-CandlestickColorsBuilder.md)
 * <span class="badge builder"></span> [CandlestickFieldMapBuilder](./builder-CandlestickFieldMapBuilder.md)
 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
## Functions

### <span class="badge function"></span> NewCandlestickFieldMap

NewCandlestickFieldMap creates a new CandlestickFieldMap object.

```go
func NewCandlestickFieldMap() *CandlestickFieldMap
```

### <span class="badge function"></span> NewCandlestickColors

NewCandlestickColors creates a new CandlestickColors object.

```go
func NewCandlestickColors() *CandlestickColors
```

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

VariantConfig returns the configuration related to candlestick panels.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.PanelcfgConfig
```

### <span class="badge function"></span> CandlestickFieldMapConverter

CandlestickFieldMapConverter accepts a `CandlestickFieldMap` object and generates the Go code to build this object using builders.

```go
func CandlestickFieldMapConverter(input CandlestickFieldMap) string
```

### <span class="badge function"></span> CandlestickColorsConverter

CandlestickColorsConverter accepts a `CandlestickColors` object and generates the Go code to build this object using builders.

```go
func CandlestickColorsConverter(input CandlestickColors) string
```

### <span class="badge function"></span> PanelConverter

PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.

```go
func PanelConverter(input dashboard.Panel) string
```

