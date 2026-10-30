---
title: <span class="badge object-type-struct"></span> RuleRef
---
# <span class="badge object-type-struct"></span> RuleRef

## Definition

```go
type RuleRef struct {
    // name is the metadata.name of an AlertRule or RecordingRule resource.
    Name alertingv1beta1.RuleUID `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (ruleRef *RuleRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RuleRef` objects.

```go
func (ruleRef *RuleRef) Equals(other RuleRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RuleRef` fields for violations and returns them.

```go
func (ruleRef *RuleRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [RuleRefBuilder](./builder-RuleRefBuilder.md)
