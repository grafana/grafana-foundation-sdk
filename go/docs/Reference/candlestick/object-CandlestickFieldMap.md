---
title: <span class="badge object-type-struct"></span> CandlestickFieldMap
---
# <span class="badge object-type-struct"></span> CandlestickFieldMap

## Definition

```go
type CandlestickFieldMap struct {
    // Corresponds to the starting value of the given period
    Open *string `json:"open,omitempty"`
    // Corresponds to the highest value of the given period
    High *string `json:"high,omitempty"`
    // Corresponds to the lowest value of the given period
    Low *string `json:"low,omitempty"`
    // Corresponds to the final (end) value of the given period
    Close *string `json:"close,omitempty"`
    // Corresponds to the sample count in the given period. (e.g. number of trades)
    Volume *string `json:"volume,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CandlestickFieldMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (candlestickFieldMap *CandlestickFieldMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CandlestickFieldMap` objects.

```go
func (candlestickFieldMap *CandlestickFieldMap) Equals(other CandlestickFieldMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CandlestickFieldMap` fields for violations and returns them.

```go
func (candlestickFieldMap *CandlestickFieldMap) Validate() error
```

## See also

 * <span class="badge builder"></span> [CandlestickFieldMapBuilder](./builder-CandlestickFieldMapBuilder.md)
