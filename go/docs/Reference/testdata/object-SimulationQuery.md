---
title: <span class="badge object-type-struct"></span> SimulationQuery
---
# <span class="badge object-type-struct"></span> SimulationQuery

## Definition

```go
type SimulationQuery struct {
    Key testdata.Key `json:"key"`
    Config map[string]any `json:"config,omitempty"`
    Stream *bool `json:"stream,omitempty"`
    Last *bool `json:"last,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SimulationQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (simulationQuery *SimulationQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SimulationQuery` objects.

```go
func (simulationQuery *SimulationQuery) Equals(other SimulationQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SimulationQuery` fields for violations and returns them.

```go
func (simulationQuery *SimulationQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [SimulationQueryBuilder](./builder-SimulationQueryBuilder.md)
