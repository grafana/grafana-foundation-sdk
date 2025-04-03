---
title: <span class="badge object-type-struct"></span> StringOrMap
---
# <span class="badge object-type-struct"></span> StringOrMap

## Definition

```go
type StringOrMap struct {
    String *string `json:"String,omitempty"`
    Map map[string]any `json:"Map,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrMap` as JSON.

```go
func (stringOrMap *StringOrMap) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrMap` from JSON.

```go
func (stringOrMap *StringOrMap) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrMap *StringOrMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrMap` objects.

```go
func (stringOrMap *StringOrMap) Equals(other StringOrMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrMap` fields for violations and returns them.

```go
func (stringOrMap *StringOrMap) Validate() error
```

