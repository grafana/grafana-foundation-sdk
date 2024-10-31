---
title: <span class="badge object-type-struct"></span> ArcOption
---
# <span class="badge object-type-struct"></span> ArcOption

## Definition

```go
type ArcOption struct {
    // Field from which to get the value. Values should be less than 1, representing fraction of a circle.
    Field *string `json:"field,omitempty"`
    // The color of the arc.
    Color *string `json:"color,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ArcOption` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (arcOption *ArcOption) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ArcOption` objects.

```go
func (arcOption *ArcOption) Equals(other ArcOption) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ArcOption` fields for violations and returns them.

```go
func (arcOption *ArcOption) Validate() error
```

## See also

 * <span class="badge builder"></span> [ArcOptionBuilder](./builder-ArcOptionBuilder.md)
