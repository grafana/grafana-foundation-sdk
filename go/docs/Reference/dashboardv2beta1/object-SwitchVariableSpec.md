---
title: <span class="badge object-type-struct"></span> SwitchVariableSpec
---
# <span class="badge object-type-struct"></span> SwitchVariableSpec

## Definition

```go
type SwitchVariableSpec struct {
    Name string `json:"name"`
    Current string `json:"current"`
    EnabledValue string `json:"enabledValue"`
    DisabledValue string `json:"disabledValue"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SwitchVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (switchVariableSpec *SwitchVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SwitchVariableSpec` objects.

```go
func (switchVariableSpec *SwitchVariableSpec) Equals(other SwitchVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SwitchVariableSpec` fields for violations and returns them.

```go
func (switchVariableSpec *SwitchVariableSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [SwitchVariableSpecBuilder](./builder-SwitchVariableSpecBuilder.md)
