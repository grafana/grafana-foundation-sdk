---
title: <span class="badge object-type-struct"></span> ValueMap
---
# <span class="badge object-type-struct"></span> ValueMap

Maps text values to a color or different display text and color.

For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.

## Definition

```go
type ValueMap struct {
    Type string `json:"type"`
    // Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
    Options map[string]dashboard.ValueMappingResult `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ValueMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (valueMap *ValueMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ValueMap` objects.

```go
func (valueMap *ValueMap) Equals(other ValueMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ValueMap` fields for violations and returns them.

```go
func (valueMap *ValueMap) Validate() error
```

## See also

 * <span class="badge builder"></span> [ValueMapBuilder](./builder-ValueMapBuilder.md)
