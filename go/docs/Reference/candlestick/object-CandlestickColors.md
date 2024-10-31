---
title: <span class="badge object-type-struct"></span> CandlestickColors
---
# <span class="badge object-type-struct"></span> CandlestickColors

## Definition

```go
type CandlestickColors struct {
    Up string `json:"up"`
    Down string `json:"down"`
    Flat string `json:"flat"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CandlestickColors` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (candlestickColors *CandlestickColors) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CandlestickColors` objects.

```go
func (candlestickColors *CandlestickColors) Equals(other CandlestickColors) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CandlestickColors` fields for violations and returns them.

```go
func (candlestickColors *CandlestickColors) Validate() error
```

## See also

 * <span class="badge builder"></span> [CandlestickColorsBuilder](./builder-CandlestickColorsBuilder.md)
