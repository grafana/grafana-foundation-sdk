---
title: <span class="badge object-type-struct"></span> V2TransformationSpec
---
# <span class="badge object-type-struct"></span> V2TransformationSpec

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Definition

```go
type V2TransformationSpec struct {
    // Disabled transformations are skipped
    Disabled *bool `json:"disabled,omitempty"`
    // Optional frame matcher. When missing it will be applied to all results
    Filter *dashboardv2beta1.MatcherConfig `json:"filter,omitempty"`
    // Where to pull DataFrames from as input to transformation
    Topic *dashboardv2beta1.DataTopic `json:"topic,omitempty"`
    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    Options any `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `V2TransformationSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (v2TransformationSpec *V2TransformationSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `V2TransformationSpec` objects.

```go
func (v2TransformationSpec *V2TransformationSpec) Equals(other V2TransformationSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `V2TransformationSpec` fields for violations and returns them.

```go
func (v2TransformationSpec *V2TransformationSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [V2TransformationSpecBuilder](./builder-V2TransformationSpecBuilder.md)
