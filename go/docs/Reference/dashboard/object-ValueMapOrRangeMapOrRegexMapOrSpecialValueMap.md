---
title: <span class="badge object-type-struct"></span> ValueMapOrRangeMapOrRegexMapOrSpecialValueMap
---
# <span class="badge object-type-struct"></span> ValueMapOrRangeMapOrRegexMapOrSpecialValueMap

## Definition

```go
type ValueMapOrRangeMapOrRegexMapOrSpecialValueMap struct {
    ValueMap *dashboard.ValueMap `json:"ValueMap,omitempty"`
    RangeMap *dashboard.RangeMap `json:"RangeMap,omitempty"`
    RegexMap *dashboard.RegexMap `json:"RegexMap,omitempty"`
    SpecialValueMap *dashboard.SpecialValueMap `json:"SpecialValueMap,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` as JSON.

```go
func (valueMapOrRangeMapOrRegexMapOrSpecialValueMap *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` from JSON.

```go
func (valueMapOrRangeMapOrRegexMapOrSpecialValueMap *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (valueMapOrRangeMapOrRegexMapOrSpecialValueMap *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` objects.

```go
func (valueMapOrRangeMapOrRegexMapOrSpecialValueMap *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) Equals(other ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` fields for violations and returns them.

```go
func (valueMapOrRangeMapOrRegexMapOrSpecialValueMap *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) Validate() error
```

