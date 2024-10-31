---
title: <span class="badge object-type-struct"></span> FiltersSettings
---
# <span class="badge object-type-struct"></span> FiltersSettings

## Definition

```go
type FiltersSettings struct {
    Filters []elasticsearch.Filter `json:"filters,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FiltersSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (filtersSettings *FiltersSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FiltersSettings` objects.

```go
func (filtersSettings *FiltersSettings) Equals(other FiltersSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FiltersSettings` fields for violations and returns them.

```go
func (filtersSettings *FiltersSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [FiltersSettingsBuilder](./builder-FiltersSettingsBuilder.md)
