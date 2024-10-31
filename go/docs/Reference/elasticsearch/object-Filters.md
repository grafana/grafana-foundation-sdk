---
title: <span class="badge object-type-struct"></span> Filters
---
# <span class="badge object-type-struct"></span> Filters

## Definition

```go
type Filters struct {
    Id string `json:"id"`
    Type string `json:"type"`
    Settings *elasticsearch.ElasticsearchFiltersSettings `json:"settings,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Filters` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (filters *Filters) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Filters` objects.

```go
func (filters *Filters) Equals(other Filters) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Filters` fields for violations and returns them.

```go
func (filters *Filters) Validate() error
```

## See also

 * <span class="badge builder"></span> [FiltersBuilder](./builder-FiltersBuilder.md)
