---
title: <span class="badge object-type-struct"></span> GroupByVariableKind
---
# <span class="badge object-type-struct"></span> GroupByVariableKind

Group variable kind

## Definition

```go
type GroupByVariableKind struct {
    Kind string `json:"kind"`
    Group string `json:"group"`
    Datasource *dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource `json:"datasource,omitempty"`
    Spec dashboardv2beta1.GroupByVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GroupByVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (groupByVariableKind *GroupByVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GroupByVariableKind` objects.

```go
func (groupByVariableKind *GroupByVariableKind) Equals(other GroupByVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GroupByVariableKind` fields for violations and returns them.

```go
func (groupByVariableKind *GroupByVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [GroupByVariableBuilder](./builder-GroupByVariableBuilder.md)
