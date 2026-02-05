---
title: <span class="badge object-type-struct"></span> XychartXYSeriesConfigColor
---
# <span class="badge object-type-struct"></span> XychartXYSeriesConfigColor

## Definition

```go
type XychartXYSeriesConfigColor struct {
    Matcher xychart.MatcherConfig `json:"matcher"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigColor` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartXYSeriesConfigColor *XychartXYSeriesConfigColor) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartXYSeriesConfigColor` objects.

```go
func (xychartXYSeriesConfigColor *XychartXYSeriesConfigColor) Equals(other XychartXYSeriesConfigColor) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigColor` fields for violations and returns them.

```go
func (xychartXYSeriesConfigColor *XychartXYSeriesConfigColor) Validate() error
```

## See also

 * <span class="badge builder"></span> [XychartXYSeriesConfigColorBuilder](./builder-XychartXYSeriesConfigColorBuilder.md)
