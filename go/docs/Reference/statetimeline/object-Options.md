---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    // Show timeline values on chart
    ShowValue common.VisibilityMode `json:"showValue"`
    // Controls the row height
    RowHeight float64 `json:"rowHeight"`
    // Merge equal consecutive values
    MergeValues *bool `json:"mergeValues,omitempty"`
    // Controls value alignment on the timelines
    AlignValue *common.TimelineValueAlignment `json:"alignValue,omitempty"`
    Legend common.VizLegendOptions `json:"legend"`
    Tooltip common.VizTooltipOptions `json:"tooltip"`
    Timezone []common.TimeZone `json:"timezone,omitempty"`
    // Enables pagination when > 0
    PerPage *float64 `json:"perPage,omitempty"`
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
