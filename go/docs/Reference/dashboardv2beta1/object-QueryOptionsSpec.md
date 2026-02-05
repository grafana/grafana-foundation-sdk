---
title: <span class="badge object-type-struct"></span> QueryOptionsSpec
---
# <span class="badge object-type-struct"></span> QueryOptionsSpec

## Definition

```go
type QueryOptionsSpec struct {
    TimeFrom *string `json:"timeFrom,omitempty"`
    MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
    TimeShift *string `json:"timeShift,omitempty"`
    QueryCachingTTL *int64 `json:"queryCachingTTL,omitempty"`
    Interval *string `json:"interval,omitempty"`
    CacheTimeout *string `json:"cacheTimeout,omitempty"`
    HideTimeOverride *bool `json:"hideTimeOverride,omitempty"`
    TimeCompare *string `json:"timeCompare,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryOptionsSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryOptionsSpec *QueryOptionsSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryOptionsSpec` objects.

```go
func (queryOptionsSpec *QueryOptionsSpec) Equals(other QueryOptionsSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryOptionsSpec` fields for violations and returns them.

```go
func (queryOptionsSpec *QueryOptionsSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryOptionsSpecBuilder](./builder-QueryOptionsSpecBuilder.md)
