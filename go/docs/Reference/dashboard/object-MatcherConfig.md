---
title: <span class="badge object-type-struct"></span> MatcherConfig
---
# <span class="badge object-type-struct"></span> MatcherConfig

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Definition

```go
type MatcherConfig struct {
    // The matcher id. This is used to find the matcher implementation from registry.
    Id string `json:"id"`
    // The matcher options. This is specific to the matcher implementation.
    Options any `json:"options,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MatcherConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (matcherConfig *MatcherConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MatcherConfig` objects.

```go
func (matcherConfig *MatcherConfig) Equals(other MatcherConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MatcherConfig` fields for violations and returns them.

```go
func (matcherConfig *MatcherConfig) Validate() error
```

