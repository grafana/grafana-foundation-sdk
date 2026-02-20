---
title: <span class="badge object-type-struct"></span> Stars
---
# <span class="badge object-type-struct"></span> Stars

## Definition

```go
type Stars struct {
    Resource []starsv1alpha1.Resource `json:"resource"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Stars` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stars *Stars) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Stars` objects.

```go
func (stars *Stars) Equals(other Stars) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Stars` fields for violations and returns them.

```go
func (stars *Stars) Validate() error
```

## See also

 * <span class="badge builder"></span> [StarsBuilder](./builder-StarsBuilder.md)
