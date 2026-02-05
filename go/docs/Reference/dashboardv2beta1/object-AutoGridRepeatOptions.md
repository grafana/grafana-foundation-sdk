---
title: <span class="badge object-type-struct"></span> AutoGridRepeatOptions
---
# <span class="badge object-type-struct"></span> AutoGridRepeatOptions

## Definition

```go
type AutoGridRepeatOptions struct {
    Mode dashboardv2beta1.RepeatMode `json:"mode"`
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridRepeatOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (autoGridRepeatOptions *AutoGridRepeatOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AutoGridRepeatOptions` objects.

```go
func (autoGridRepeatOptions *AutoGridRepeatOptions) Equals(other AutoGridRepeatOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AutoGridRepeatOptions` fields for violations and returns them.

```go
func (autoGridRepeatOptions *AutoGridRepeatOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [AutoGridRepeatOptionsBuilder](./builder-AutoGridRepeatOptionsBuilder.md)
