---
title: <span class="badge object-type-struct"></span> RuleGroup
---
# <span class="badge object-type-struct"></span> RuleGroup

## Definition

```go
type RuleGroup struct {
    FolderUid *string `json:"folderUid,omitempty"`
    // The interval, in seconds, at which all rules in the group are evaluated.
    // If a group contains many rules, the rules are evaluated sequentially.
    Interval *alerting.Duration `json:"interval,omitempty"`
    Rules []alerting.Rule `json:"rules,omitempty"`
    Title *string `json:"title,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleGroup` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (ruleGroup *RuleGroup) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RuleGroup` objects.

```go
func (ruleGroup *RuleGroup) Equals(other RuleGroup) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RuleGroup` fields for violations and returns them.

```go
func (ruleGroup *RuleGroup) Validate() error
```

## See also

 * <span class="badge builder"></span> [RuleGroupBuilder](./builder-RuleGroupBuilder.md)
