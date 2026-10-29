---
title: <span class="badge object-type-struct"></span> RuleSpec
---
# <span class="badge object-type-struct"></span> RuleSpec

## Definition

```go
type RuleSpec struct {
    Title string `json:"title"`
    Paused *bool `json:"paused,omitempty"`
    Trigger alertingv0alpha1.IntervalTrigger `json:"trigger"`
    Labels map[string]alertingv0alpha1.TemplateString `json:"labels,omitempty"`
    Expressions alertingv0alpha1.ExpressionMap `json:"expressions"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (ruleSpec *RuleSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RuleSpec` objects.

```go
func (ruleSpec *RuleSpec) Equals(other RuleSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RuleSpec` fields for violations and returns them.

```go
func (ruleSpec *RuleSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [RuleSpecBuilder](./builder-RuleSpecBuilder.md)
