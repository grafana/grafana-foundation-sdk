---
title: <span class="badge object-type-struct"></span> Dashboardv2SpecialValueMapOptions
---
# <span class="badge object-type-struct"></span> Dashboardv2SpecialValueMapOptions

## Definition

```go
type Dashboardv2SpecialValueMapOptions struct {
    // Special value to match against
    Match dashboardv2.SpecialValueMatch `json:"match"`
    // Config to apply when the value matches the special value
    Result dashboardv2.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2SpecialValueMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2SpecialValueMapOptions *Dashboardv2SpecialValueMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2SpecialValueMapOptions` objects.

```go
func (dashboardv2SpecialValueMapOptions *Dashboardv2SpecialValueMapOptions) Equals(other Dashboardv2SpecialValueMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2SpecialValueMapOptions` fields for violations and returns them.

```go
func (dashboardv2SpecialValueMapOptions *Dashboardv2SpecialValueMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2SpecialValueMapOptionsBuilder](./builder-Dashboardv2SpecialValueMapOptionsBuilder.md)
