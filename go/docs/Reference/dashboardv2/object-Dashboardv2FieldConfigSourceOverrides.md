---
title: <span class="badge object-type-struct"></span> Dashboardv2FieldConfigSourceOverrides
---
# <span class="badge object-type-struct"></span> Dashboardv2FieldConfigSourceOverrides

## Definition

```go
type Dashboardv2FieldConfigSourceOverrides struct {
    SystemRef *string `json:"__systemRef,omitempty"`
    Matcher dashboardv2.MatcherConfig `json:"matcher"`
    Properties []dashboardv2.DynamicConfigValue `json:"properties"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2FieldConfigSourceOverrides` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2FieldConfigSourceOverrides *Dashboardv2FieldConfigSourceOverrides) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2FieldConfigSourceOverrides` objects.

```go
func (dashboardv2FieldConfigSourceOverrides *Dashboardv2FieldConfigSourceOverrides) Equals(other Dashboardv2FieldConfigSourceOverrides) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2FieldConfigSourceOverrides` fields for violations and returns them.

```go
func (dashboardv2FieldConfigSourceOverrides *Dashboardv2FieldConfigSourceOverrides) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2FieldConfigSourceOverridesBuilder](./builder-Dashboardv2FieldConfigSourceOverridesBuilder.md)
