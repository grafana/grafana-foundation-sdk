---
title: <span class="badge object-type-struct"></span> Rule
---
# <span class="badge object-type-struct"></span> Rule

## Definition

```go
type Rule struct {
    Annotations map[string]string `json:"annotations,omitempty"`
    Condition string `json:"condition"`
    Data []alerting.Query `json:"data"`
    ExecErrState alerting.RuleExecErrState `json:"execErrState"`
    FolderUID string `json:"folderUID"`
    // The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    // Before this time has elapsed, the rule is only considered to be Pending.
    For string `json:"for"`
    Id *int64 `json:"id,omitempty"`
    IsPaused *bool `json:"isPaused,omitempty"`
    Labels map[string]string `json:"labels,omitempty"`
    NoDataState alerting.RuleNoDataState `json:"noDataState"`
    OrgID int64 `json:"orgID"`
    Provenance *alerting.Provenance `json:"provenance,omitempty"`
    RuleGroup string `json:"ruleGroup"`
    Title string `json:"title"`
    Uid *string `json:"uid,omitempty"`
    Updated *time.Time `json:"updated,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Rule` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rule *Rule) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Rule` objects.

```go
func (rule *Rule) Equals(other Rule) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Rule` fields for violations and returns them.

```go
func (rule *Rule) Validate() error
```

## See also

 * <span class="badge builder"></span> [RuleBuilder](./builder-RuleBuilder.md)
