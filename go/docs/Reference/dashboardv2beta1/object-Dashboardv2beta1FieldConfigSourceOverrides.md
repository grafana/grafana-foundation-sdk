---
title: <span class="badge object-type-struct"></span> Dashboardv2beta1FieldConfigSourceOverrides
---
# <span class="badge object-type-struct"></span> Dashboardv2beta1FieldConfigSourceOverrides

## Definition

```go
type Dashboardv2beta1FieldConfigSourceOverrides struct {
    SystemRef *string `json:"__systemRef,omitempty"`
    Matcher dashboardv2beta1.MatcherConfig `json:"matcher"`
    Properties []dashboardv2beta1.DynamicConfigValue `json:"properties"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1FieldConfigSourceOverrides` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2beta1FieldConfigSourceOverrides *Dashboardv2beta1FieldConfigSourceOverrides) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2beta1FieldConfigSourceOverrides` objects.

```go
func (dashboardv2beta1FieldConfigSourceOverrides *Dashboardv2beta1FieldConfigSourceOverrides) Equals(other Dashboardv2beta1FieldConfigSourceOverrides) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2beta1FieldConfigSourceOverrides` fields for violations and returns them.

```go
func (dashboardv2beta1FieldConfigSourceOverrides *Dashboardv2beta1FieldConfigSourceOverrides) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2beta1FieldConfigSourceOverridesBuilder](./builder-Dashboardv2beta1FieldConfigSourceOverridesBuilder.md)
