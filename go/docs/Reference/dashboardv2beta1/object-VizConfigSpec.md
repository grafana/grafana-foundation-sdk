---
title: <span class="badge object-type-struct"></span> VizConfigSpec
---
# <span class="badge object-type-struct"></span> VizConfigSpec

--- Kinds ---

## Definition

```go
type VizConfigSpec struct {
    Options any `json:"options"`
    FieldConfig dashboardv2beta1.FieldConfigSource `json:"fieldConfig"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizConfigSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (vizConfigSpec *VizConfigSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VizConfigSpec` objects.

```go
func (vizConfigSpec *VizConfigSpec) Equals(other VizConfigSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VizConfigSpec` fields for violations and returns them.

```go
func (vizConfigSpec *VizConfigSpec) Validate() error
```

