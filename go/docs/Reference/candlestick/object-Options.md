---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    // Sets which dimensions are used for the visualization
    Mode candlestick.VizDisplayMode `json:"mode"`
    // Sets the style of the candlesticks
    CandleStyle candlestick.CandleStyle `json:"candleStyle"`
    // Sets the color strategy for the candlesticks
    ColorStrategy candlestick.ColorStrategy `json:"colorStrategy"`
    // Map fields to appropriate dimension
    Fields candlestick.CandlestickFieldMap `json:"fields"`
    // Set which colors are used when the price movement is up or down
    Colors candlestick.CandlestickColors `json:"colors"`
    Legend common.VizLegendOptions `json:"legend"`
    // When enabled, all fields will be sent to the graph
    IncludeAllFields *bool `json:"includeAllFields,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (options *Options) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Options` objects.

```go
func (options *Options) Equals(other Options) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.

```go
func (options *Options) Validate() error
```

