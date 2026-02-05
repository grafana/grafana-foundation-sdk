---
title: <span class="badge object-type-struct"></span> RowRepeatOptions
---
# <span class="badge object-type-struct"></span> RowRepeatOptions

## Definition

```go
type RowRepeatOptions struct {
    Mode dashboardv2beta1.RepeatMode `json:"mode"`
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowRepeatOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rowRepeatOptions *RowRepeatOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RowRepeatOptions` objects.

```go
func (rowRepeatOptions *RowRepeatOptions) Equals(other RowRepeatOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RowRepeatOptions` fields for violations and returns them.

```go
func (rowRepeatOptions *RowRepeatOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [RowRepeatOptionsBuilder](./builder-RowRepeatOptionsBuilder.md)
