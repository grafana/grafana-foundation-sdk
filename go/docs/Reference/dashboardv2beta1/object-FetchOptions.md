---
title: <span class="badge object-type-struct"></span> FetchOptions
---
# <span class="badge object-type-struct"></span> FetchOptions

## Definition

```go
type FetchOptions struct {
    Method dashboardv2beta1.HttpRequestMethod `json:"method"`
    Url string `json:"url"`
    Body *string `json:"body,omitempty"`
    // These are 2D arrays of strings, each representing a key-value pair
    // We are defining them this way because we can't generate a go struct that
    // that would have exactly two strings in each sub-array
    QueryParams [][]string `json:"queryParams,omitempty"`
    Headers [][]string `json:"headers,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FetchOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fetchOptions *FetchOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FetchOptions` objects.

```go
func (fetchOptions *FetchOptions) Equals(other FetchOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FetchOptions` fields for violations and returns them.

```go
func (fetchOptions *FetchOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [FetchOptionsBuilder](./builder-FetchOptionsBuilder.md)
