---
title: <span class="badge object-type-struct"></span> TimeOption
---
# <span class="badge object-type-struct"></span> TimeOption

Counterpart for TypeScript's TimeOption type.

## Definition

```go
type TimeOption struct {
    Display string `json:"display"`
    From string `json:"from"`
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeOption` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeOption *TimeOption) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeOption` objects.

```go
func (timeOption *TimeOption) Equals(other TimeOption) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeOption` fields for violations and returns them.

```go
func (timeOption *TimeOption) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeOptionBuilder](./builder-TimeOptionBuilder.md)
