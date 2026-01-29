---
title: <span class="badge object-type-struct"></span> XychartXYSeriesConfigY
---
# <span class="badge object-type-struct"></span> XychartXYSeriesConfigY

## Definition

```go
type XychartXYSeriesConfigY struct {
    Matcher xychart.MatcherConfig `json:"matcher"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigY` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartXYSeriesConfigY *XychartXYSeriesConfigY) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartXYSeriesConfigY` objects.

```go
func (xychartXYSeriesConfigY *XychartXYSeriesConfigY) Equals(other XychartXYSeriesConfigY) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigY` fields for violations and returns them.

```go
func (xychartXYSeriesConfigY *XychartXYSeriesConfigY) Validate() error
```

