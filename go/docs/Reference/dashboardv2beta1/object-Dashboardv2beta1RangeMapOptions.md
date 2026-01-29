---
title: <span class="badge object-type-struct"></span> Dashboardv2beta1RangeMapOptions
---
# <span class="badge object-type-struct"></span> Dashboardv2beta1RangeMapOptions

## Definition

```go
type Dashboardv2beta1RangeMapOptions struct {
    // Min value of the range. It can be null which means -Infinity
    From *float64 `json:"from"`
    // Max value of the range. It can be null which means +Infinity
    To *float64 `json:"to"`
    // Config to apply when the value is within the range
    Result dashboardv2beta1.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1RangeMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2beta1RangeMapOptions *Dashboardv2beta1RangeMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2beta1RangeMapOptions` objects.

```go
func (dashboardv2beta1RangeMapOptions *Dashboardv2beta1RangeMapOptions) Equals(other Dashboardv2beta1RangeMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2beta1RangeMapOptions` fields for violations and returns them.

```go
func (dashboardv2beta1RangeMapOptions *Dashboardv2beta1RangeMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2beta1RangeMapOptionsBuilder](./builder-Dashboardv2beta1RangeMapOptionsBuilder.md)
