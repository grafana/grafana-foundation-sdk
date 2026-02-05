---
title: <span class="badge object-type-struct"></span> Action
---
# <span class="badge object-type-struct"></span> Action

## Definition

```go
type Action struct {
    Type dashboardv2beta1.ActionType `json:"type"`
    Title string `json:"title"`
    Fetch *dashboardv2beta1.FetchOptions `json:"fetch,omitempty"`
    Infinity *dashboardv2beta1.InfinityOptions `json:"infinity,omitempty"`
    Confirmation *string `json:"confirmation,omitempty"`
    OneClick *bool `json:"oneClick,omitempty"`
    Variables []dashboardv2beta1.ActionVariable `json:"variables,omitempty"`
    Style *dashboardv2beta1.Dashboardv2beta1ActionStyle `json:"style,omitempty"`
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
