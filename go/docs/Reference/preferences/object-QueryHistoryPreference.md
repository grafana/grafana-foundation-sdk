---
title: <span class="badge object-type-struct"></span> QueryHistoryPreference
---
# <span class="badge object-type-struct"></span> QueryHistoryPreference

## Definition

```go
type QueryHistoryPreference struct {
    // one of: '' | 'query' | 'starred';
    HomeTab *string `json:"homeTab,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryHistoryPreference` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryHistoryPreference *QueryHistoryPreference) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryHistoryPreference` objects.

```go
func (queryHistoryPreference *QueryHistoryPreference) Equals(other QueryHistoryPreference) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryHistoryPreference` fields for violations and returns them.

```go
func (queryHistoryPreference *QueryHistoryPreference) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryHistoryPreferenceBuilder](./builder-QueryHistoryPreferenceBuilder.md)
