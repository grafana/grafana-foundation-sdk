---
title: <span class="badge object-type-struct"></span> Manifest
---
# <span class="badge object-type-struct"></span> Manifest

## Definition

```go
type Manifest struct {
    ApiVersion string `json:"apiVersion"`
    Kind string `json:"kind"`
    Metadata resource.Metadata `json:"metadata"`
    Spec any `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Manifest` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (manifest *Manifest) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Manifest` objects.

```go
func (manifest *Manifest) Equals(other Manifest) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Manifest` fields for violations and returns them.

```go
func (manifest *Manifest) Validate() error
```

## See also

 * <span class="badge builder"></span> [ManifestBuilder](./builder-ManifestBuilder.md)
