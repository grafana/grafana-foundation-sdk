---
title: <span class="badge object-type-struct"></span> XychartXYSeriesConfigX
---
# <span class="badge object-type-struct"></span> XychartXYSeriesConfigX

## Definition

```go
type XychartXYSeriesConfigX struct {
    Matcher xychart.MatcherConfig `json:"matcher"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigX` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartXYSeriesConfigX *XychartXYSeriesConfigX) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartXYSeriesConfigX` objects.

```go
func (xychartXYSeriesConfigX *XychartXYSeriesConfigX) Equals(other XychartXYSeriesConfigX) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigX` fields for violations and returns them.

```go
func (xychartXYSeriesConfigX *XychartXYSeriesConfigX) Validate() error
```

