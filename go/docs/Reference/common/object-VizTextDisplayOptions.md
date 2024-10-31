---
title: <span class="badge object-type-struct"></span> VizTextDisplayOptions
---
# <span class="badge object-type-struct"></span> VizTextDisplayOptions

TODO docs

## Definition

```go
type VizTextDisplayOptions struct {
    // Explicit title text size
    TitleSize *float64 `json:"titleSize,omitempty"`
    // Explicit value text size
    ValueSize *float64 `json:"valueSize,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizTextDisplayOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (vizTextDisplayOptions *VizTextDisplayOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VizTextDisplayOptions` objects.

```go
func (vizTextDisplayOptions *VizTextDisplayOptions) Equals(other VizTextDisplayOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VizTextDisplayOptions` fields for violations and returns them.

```go
func (vizTextDisplayOptions *VizTextDisplayOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [VizTextDisplayOptionsBuilder](./builder-VizTextDisplayOptionsBuilder.md)
