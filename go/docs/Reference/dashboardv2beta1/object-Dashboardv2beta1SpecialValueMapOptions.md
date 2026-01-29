---
title: <span class="badge object-type-struct"></span> Dashboardv2beta1SpecialValueMapOptions
---
# <span class="badge object-type-struct"></span> Dashboardv2beta1SpecialValueMapOptions

## Definition

```go
type Dashboardv2beta1SpecialValueMapOptions struct {
    // Special value to match against
    Match dashboardv2beta1.SpecialValueMatch `json:"match"`
    // Config to apply when the value matches the special value
    Result dashboardv2beta1.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1SpecialValueMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2beta1SpecialValueMapOptions *Dashboardv2beta1SpecialValueMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2beta1SpecialValueMapOptions` objects.

```go
func (dashboardv2beta1SpecialValueMapOptions *Dashboardv2beta1SpecialValueMapOptions) Equals(other Dashboardv2beta1SpecialValueMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2beta1SpecialValueMapOptions` fields for violations and returns them.

```go
func (dashboardv2beta1SpecialValueMapOptions *Dashboardv2beta1SpecialValueMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2beta1SpecialValueMapOptionsBuilder](./builder-Dashboardv2beta1SpecialValueMapOptionsBuilder.md)
