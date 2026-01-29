---
title: <span class="badge object-type-struct"></span> AdHocFilterWithLabels
---
# <span class="badge object-type-struct"></span> AdHocFilterWithLabels

Define the AdHocFilterWithLabels type

## Definition

```go
type AdHocFilterWithLabels struct {
    Key string `json:"key"`
    Operator string `json:"operator"`
    Value string `json:"value"`
    Values []string `json:"values,omitempty"`
    KeyLabel *string `json:"keyLabel,omitempty"`
    ValueLabels []string `json:"valueLabels,omitempty"`
    ForceEdit *bool `json:"forceEdit,omitempty"`
    Origin *string `json:"origin,omitempty"`
    // @deprecated
    Condition *string `json:"condition,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdHocFilterWithLabels` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (adHocFilterWithLabels *AdHocFilterWithLabels) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AdHocFilterWithLabels` objects.

```go
func (adHocFilterWithLabels *AdHocFilterWithLabels) Equals(other AdHocFilterWithLabels) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AdHocFilterWithLabels` fields for violations and returns them.

```go
func (adHocFilterWithLabels *AdHocFilterWithLabels) Validate() error
```

## See also

 * <span class="badge builder"></span> [AdHocFilterWithLabelsBuilder](./builder-AdHocFilterWithLabelsBuilder.md)
