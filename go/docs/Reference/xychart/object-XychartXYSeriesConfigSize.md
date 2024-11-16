---
title: <span class="badge object-type-struct"></span> XychartXYSeriesConfigSize
---
# <span class="badge object-type-struct"></span> XychartXYSeriesConfigSize

## Definition

```go
type XychartXYSeriesConfigSize struct {
    Matcher xychart.MatcherConfig `json:"matcher"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigSize` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartXYSeriesConfigSize *XychartXYSeriesConfigSize) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartXYSeriesConfigSize` objects.

```go
func (xychartXYSeriesConfigSize *XychartXYSeriesConfigSize) Equals(other XychartXYSeriesConfigSize) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigSize` fields for violations and returns them.

```go
func (xychartXYSeriesConfigSize *XychartXYSeriesConfigSize) Validate() error
```

## See also

 * <span class="badge builder"></span> [XychartXYSeriesConfigSizeBuilder](./builder-XychartXYSeriesConfigSizeBuilder.md)
