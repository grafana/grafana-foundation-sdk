---
title: <span class="badge object-type-struct"></span> AdhocVariableKind
---
# <span class="badge object-type-struct"></span> AdhocVariableKind

Adhoc variable kind

## Definition

```go
type AdhocVariableKind struct {
    Kind string `json:"kind"`
    Group string `json:"group"`
    Datasource *dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource `json:"datasource,omitempty"`
    Spec dashboardv2beta1.AdhocVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdhocVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (adhocVariableKind *AdhocVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AdhocVariableKind` objects.

```go
func (adhocVariableKind *AdhocVariableKind) Equals(other AdhocVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AdhocVariableKind` fields for violations and returns them.

```go
func (adhocVariableKind *AdhocVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [AdhocVariableBuilder](./builder-AdhocVariableBuilder.md)
