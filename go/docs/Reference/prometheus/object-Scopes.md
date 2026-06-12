---
title: <span class="badge object-type-struct"></span> Scopes
---
# <span class="badge object-type-struct"></span> Scopes

## Definition

```go
type Scopes struct {
    DefaultPath []string `json:"defaultPath,omitempty"`
    Filters []prometheus.ScopesFilters `json:"filters,omitempty"`
    Title string `json:"title"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Scopes` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (scopes *Scopes) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Scopes` objects.

```go
func (scopes *Scopes) Equals(other Scopes) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Scopes` fields for violations and returns them.

```go
func (scopes *Scopes) Validate() error
```

## See also

 * <span class="badge builder"></span> [ScopesBuilder](./builder-ScopesBuilder.md)
