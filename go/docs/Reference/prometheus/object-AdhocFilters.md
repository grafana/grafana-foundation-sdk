---
title: <span class="badge object-type-struct"></span> AdhocFilters
---
# <span class="badge object-type-struct"></span> AdhocFilters

## Definition

```go
type AdhocFilters struct {
    Key string `json:"key"`
    Operator string `json:"operator"`
    Value string `json:"value"`
    Values []string `json:"values,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdhocFilters` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (adhocFilters *AdhocFilters) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AdhocFilters` objects.

```go
func (adhocFilters *AdhocFilters) Equals(other AdhocFilters) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AdhocFilters` fields for violations and returns them.

```go
func (adhocFilters *AdhocFilters) Validate() error
```

## See also

 * <span class="badge builder"></span> [AdhocFiltersBuilder](./builder-AdhocFiltersBuilder.md)
