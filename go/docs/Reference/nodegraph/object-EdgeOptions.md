---
title: <span class="badge object-type-struct"></span> EdgeOptions
---
# <span class="badge object-type-struct"></span> EdgeOptions

## Definition

```go
type EdgeOptions struct {
    // Unit for the main stat to override what ever is set in the data frame.
    MainStatUnit *string `json:"mainStatUnit,omitempty"`
    // Unit for the secondary stat to override what ever is set in the data frame.
    SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `EdgeOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (edgeOptions *EdgeOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `EdgeOptions` objects.

```go
func (edgeOptions *EdgeOptions) Equals(other EdgeOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `EdgeOptions` fields for violations and returns them.

```go
func (edgeOptions *EdgeOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [EdgeOptionsBuilder](./builder-EdgeOptionsBuilder.md)
