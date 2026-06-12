---
title: <span class="badge object-type-struct"></span> ScopesFilters
---
# <span class="badge object-type-struct"></span> ScopesFilters

## Definition

```go
type ScopesFilters struct {
    Key string `json:"key"`
    Operator string `json:"operator"`
    Value string `json:"value"`
    Values []string `json:"values,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScopesFilters` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (scopesFilters *ScopesFilters) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ScopesFilters` objects.

```go
func (scopesFilters *ScopesFilters) Equals(other ScopesFilters) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ScopesFilters` fields for violations and returns them.

```go
func (scopesFilters *ScopesFilters) Validate() error
```

## See also

 * <span class="badge builder"></span> [ScopesFiltersBuilder](./builder-ScopesFiltersBuilder.md)
