---
title: <span class="badge object-type-struct"></span> XychartXYSeriesConfigFrame
---
# <span class="badge object-type-struct"></span> XychartXYSeriesConfigFrame

## Definition

```go
type XychartXYSeriesConfigFrame struct {
    Matcher xychart.MatcherConfig `json:"matcher"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigFrame` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartXYSeriesConfigFrame *XychartXYSeriesConfigFrame) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartXYSeriesConfigFrame` objects.

```go
func (xychartXYSeriesConfigFrame *XychartXYSeriesConfigFrame) Equals(other XychartXYSeriesConfigFrame) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigFrame` fields for violations and returns them.

```go
func (xychartXYSeriesConfigFrame *XychartXYSeriesConfigFrame) Validate() error
```

