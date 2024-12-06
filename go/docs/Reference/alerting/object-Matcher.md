---
title: <span class="badge object-type-struct"></span> Matcher
---
# <span class="badge object-type-struct"></span> Matcher

## Definition

```go
type Matcher struct {
    Name *string `json:"Name,omitempty"`
    Type *alerting.MatchType `json:"Type,omitempty"`
    Value *string `json:"Value,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Matcher` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (matcher *Matcher) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Matcher` objects.

```go
func (matcher *Matcher) Equals(other Matcher) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Matcher` fields for violations and returns them.

```go
func (matcher *Matcher) Validate() error
```

## See also

 * <span class="badge builder"></span> [MatcherBuilder](./builder-MatcherBuilder.md)
