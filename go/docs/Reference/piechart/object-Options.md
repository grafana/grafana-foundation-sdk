---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    PieType piechart.PieChartType `json:"pieType"`
    DisplayLabels []piechart.PieChartLabels `json:"displayLabels,omitempty"`
    Tooltip common.VizTooltipOptions `json:"tooltip"`
    ReduceOptions common.ReduceDataOptions `json:"reduceOptions"`
    Text *common.VizTextDisplayOptions `json:"text,omitempty"`
    Legend piechart.PieChartLegendOptions `json:"legend"`
    Orientation common.VizOrientation `json:"orientation"`
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

