---
title: <span class="badge object-type-struct"></span> XYSeriesConfig
---
# <span class="badge object-type-struct"></span> XYSeriesConfig

## Definition

```go
type XYSeriesConfig struct {
    Name *xychart.XychartXYSeriesConfigName `json:"name,omitempty"`
    Frame *xychart.XychartXYSeriesConfigFrame `json:"frame,omitempty"`
    X *xychart.XychartXYSeriesConfigX `json:"x,omitempty"`
    Y *xychart.XychartXYSeriesConfigY `json:"y,omitempty"`
    Color *xychart.XychartXYSeriesConfigColor `json:"color,omitempty"`
    Size *xychart.XychartXYSeriesConfigSize `json:"size,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XYSeriesConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xYSeriesConfig *XYSeriesConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XYSeriesConfig` objects.

```go
func (xYSeriesConfig *XYSeriesConfig) Equals(other XYSeriesConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XYSeriesConfig` fields for violations and returns them.

```go
func (xYSeriesConfig *XYSeriesConfig) Validate() error
```

