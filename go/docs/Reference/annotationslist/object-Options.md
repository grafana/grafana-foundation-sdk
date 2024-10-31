---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    OnlyFromThisDashboard bool `json:"onlyFromThisDashboard"`
    OnlyInTimeRange bool `json:"onlyInTimeRange"`
    Tags []string `json:"tags"`
    Limit uint32 `json:"limit"`
    ShowUser bool `json:"showUser"`
    ShowTime bool `json:"showTime"`
    ShowTags bool `json:"showTags"`
    NavigateToPanel bool `json:"navigateToPanel"`
    NavigateBefore string `json:"navigateBefore"`
    NavigateAfter string `json:"navigateAfter"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (options *Options) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Options` objects.

```go
func (options *Options) Equals(other Options) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.

```go
func (options *Options) Validate() error
```

