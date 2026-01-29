---
title: <span class="badge object-type-struct"></span> CustomFormatterVariable
---
# <span class="badge object-type-struct"></span> CustomFormatterVariable

Custom formatter variable

## Definition

```go
type CustomFormatterVariable struct {
    Name string `json:"name"`
    Type dashboardv2beta1.VariableType `json:"type"`
    Multi bool `json:"multi"`
    IncludeAll bool `json:"includeAll"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomFormatterVariable` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (customFormatterVariable *CustomFormatterVariable) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CustomFormatterVariable` objects.

```go
func (customFormatterVariable *CustomFormatterVariable) Equals(other CustomFormatterVariable) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CustomFormatterVariable` fields for violations and returns them.

```go
func (customFormatterVariable *CustomFormatterVariable) Validate() error
```

## See also

 * <span class="badge builder"></span> [CustomFormatterVariableBuilder](./builder-CustomFormatterVariableBuilder.md)
