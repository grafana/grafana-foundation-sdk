---
title: <span class="badge object-type-struct"></span> SpecialValueMap
---
# <span class="badge object-type-struct"></span> SpecialValueMap

Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.

See SpecialValueMatch to see the list of special values.

For example, you can configure a special value mapping so that null values appear as N/A.

## Definition

```go
type SpecialValueMap struct {
    Type string `json:"type"`
    Options dashboard.DashboardSpecialValueMapOptions `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SpecialValueMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (specialValueMap *SpecialValueMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SpecialValueMap` objects.

```go
func (specialValueMap *SpecialValueMap) Equals(other SpecialValueMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SpecialValueMap` fields for violations and returns them.

```go
func (specialValueMap *SpecialValueMap) Validate() error
```

## See also

 * <span class="badge builder"></span> [SpecialValueMapBuilder](./builder-SpecialValueMapBuilder.md)
