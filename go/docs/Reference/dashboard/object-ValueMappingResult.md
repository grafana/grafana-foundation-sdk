---
title: <span class="badge object-type-struct"></span> ValueMappingResult
---
# <span class="badge object-type-struct"></span> ValueMappingResult

Result used as replacement with text and color when the value matches

## Definition

```go
type ValueMappingResult struct {
    // Text to display when the value matches
    Text *string `json:"text,omitempty"`
    // Text to use when the value matches
    Color *string `json:"color,omitempty"`
    // Icon to display when the value matches. Only specific visualizations.
    Icon *string `json:"icon,omitempty"`
    // Position in the mapping array. Only used internally.
    Index *int32 `json:"index,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ValueMappingResult` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (valueMappingResult *ValueMappingResult) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ValueMappingResult` objects.

```go
func (valueMappingResult *ValueMappingResult) Equals(other ValueMappingResult) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ValueMappingResult` fields for violations and returns them.

```go
func (valueMappingResult *ValueMappingResult) Validate() error
```

## See also

 * <span class="badge builder"></span> [ValueMappingResultBuilder](./builder-ValueMappingResultBuilder.md)
