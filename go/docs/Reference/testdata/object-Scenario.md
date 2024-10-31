---
title: <span class="badge object-type-struct"></span> Scenario
---
# <span class="badge object-type-struct"></span> Scenario

TODO: Should this live here given it's not used in the dataquery?

## Definition

```go
type Scenario struct {
    Id string `json:"id"`
    Name string `json:"name"`
    StringInput string `json:"stringInput"`
    Description *string `json:"description,omitempty"`
    HideAliasField *bool `json:"hideAliasField,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Scenario` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (scenario *Scenario) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Scenario` objects.

```go
func (scenario *Scenario) Equals(other Scenario) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Scenario` fields for violations and returns them.

```go
func (scenario *Scenario) Validate() error
```

## See also

 * <span class="badge builder"></span> [ScenarioBuilder](./builder-ScenarioBuilder.md)
