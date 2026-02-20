---
title: <span class="badge object-type-struct"></span> Action
---
# <span class="badge object-type-struct"></span> Action

Dashboard action

## Definition

```go
type Action struct {
    Type dashboard.ActionType `json:"type"`
    Title string `json:"title"`
    Fetch *dashboard.FetchOptions `json:"fetch,omitempty"`
    Infinity *dashboard.InfinityOptions `json:"infinity,omitempty"`
    Confirmation *string `json:"confirmation,omitempty"`
    OneClick *bool `json:"oneClick,omitempty"`
    Variables []dashboard.ActionVariable `json:"variables,omitempty"`
    Style *dashboard.DashboardActionStyle `json:"style,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Action` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (action *Action) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Action` objects.

```go
func (action *Action) Equals(other Action) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Action` fields for violations and returns them.

```go
func (action *Action) Validate() error
```

## See also

 * <span class="badge builder"></span> [ActionBuilder](./builder-ActionBuilder.md)
