---
title: <span class="badge object-type-struct"></span> XychartFieldConfigPointSize
---
# <span class="badge object-type-struct"></span> XychartFieldConfigPointSize

## Definition

```go
type XychartFieldConfigPointSize struct {
    Fixed *int32 `json:"fixed,omitempty"`
    Min *int32 `json:"min,omitempty"`
    Max *int32 `json:"max,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartFieldConfigPointSize` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartFieldConfigPointSize *XychartFieldConfigPointSize) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartFieldConfigPointSize` objects.

```go
func (xychartFieldConfigPointSize *XychartFieldConfigPointSize) Equals(other XychartFieldConfigPointSize) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartFieldConfigPointSize` fields for violations and returns them.

```go
func (xychartFieldConfigPointSize *XychartFieldConfigPointSize) Validate() error
```

## See also

 * <span class="badge builder"></span> [XychartFieldConfigPointSizeBuilder](./builder-XychartFieldConfigPointSizeBuilder.md)
