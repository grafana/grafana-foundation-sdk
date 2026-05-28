---
title: <span class="badge object-type-struct"></span> DynamicConfigValue
---
# <span class="badge object-type-struct"></span> DynamicConfigValue

## Definition

```go
type DynamicConfigValue struct {
    Id string `json:"id"`
    Value any `json:"value,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DynamicConfigValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dynamicConfigValue *DynamicConfigValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DynamicConfigValue` objects.

```go
func (dynamicConfigValue *DynamicConfigValue) Equals(other DynamicConfigValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DynamicConfigValue` fields for violations and returns them.

```go
func (dynamicConfigValue *DynamicConfigValue) Validate() error
```

