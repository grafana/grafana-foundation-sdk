---
title: <span class="badge object-type-struct"></span> StackableFieldConfig
---
# <span class="badge object-type-struct"></span> StackableFieldConfig

TODO docs

## Definition

```go
type StackableFieldConfig struct {
    Stacking *common.StackingConfig `json:"stacking,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StackableFieldConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stackableFieldConfig *StackableFieldConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StackableFieldConfig` objects.

```go
func (stackableFieldConfig *StackableFieldConfig) Equals(other StackableFieldConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StackableFieldConfig` fields for violations and returns them.

```go
func (stackableFieldConfig *StackableFieldConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [StackableFieldConfigBuilder](./builder-StackableFieldConfigBuilder.md)
