---
title: <span class="badge object-type-struct"></span> ActionVariable
---
# <span class="badge object-type-struct"></span> ActionVariable

## Definition

```go
type ActionVariable struct {
    Key string `json:"key"`
    Name string `json:"name"`
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ActionVariable` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (actionVariable *ActionVariable) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ActionVariable` objects.

```go
func (actionVariable *ActionVariable) Equals(other ActionVariable) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ActionVariable` fields for violations and returns them.

```go
func (actionVariable *ActionVariable) Validate() error
```

## See also

 * <span class="badge builder"></span> [ActionVariableBuilder](./builder-ActionVariableBuilder.md)
