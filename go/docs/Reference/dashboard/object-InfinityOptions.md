---
title: <span class="badge object-type-struct"></span> InfinityOptions
---
# <span class="badge object-type-struct"></span> InfinityOptions

Infinity options

## Definition

```go
type InfinityOptions struct {
    Method dashboard.HttpRequestMethod `json:"method"`
    Url string `json:"url"`
    Body *string `json:"body,omitempty"`
    // These are 2D arrays of strings, each representing a key-value pair
    // We are defining them this way because we can't generate a go struct that
    // that would have exactly two strings in each sub-array
    QueryParams [][]string `json:"queryParams,omitempty"`
    Headers [][]string `json:"headers,omitempty"`
    DatasourceUid string `json:"datasourceUid"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `InfinityOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (infinityOptions *InfinityOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `InfinityOptions` objects.

```go
func (infinityOptions *InfinityOptions) Equals(other InfinityOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `InfinityOptions` fields for violations and returns them.

```go
func (infinityOptions *InfinityOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [InfinityOptionsBuilder](./builder-InfinityOptionsBuilder.md)
