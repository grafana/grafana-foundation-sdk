---
title: <span class="badge object-type-struct"></span> SelectableValue
---
# <span class="badge object-type-struct"></span> SelectableValue

## Definition

```go
type SelectableValue struct {
    Label string `json:"label"`
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SelectableValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (selectableValue *SelectableValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SelectableValue` objects.

```go
func (selectableValue *SelectableValue) Equals(other SelectableValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SelectableValue` fields for violations and returns them.

```go
func (selectableValue *SelectableValue) Validate() error
```

## See also

 * <span class="badge builder"></span> [SelectableValueBuilder](./builder-SelectableValueBuilder.md)
