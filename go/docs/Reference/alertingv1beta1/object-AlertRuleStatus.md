---
title: <span class="badge object-type-struct"></span> AlertRuleStatus
---
# <span class="badge object-type-struct"></span> AlertRuleStatus

## Definition

```go
type AlertRuleStatus struct {
    Health *alertingv1beta1.AlertRuleHealth `json:"health,omitempty"`
    State *alertingv1beta1.AlertRuleState `json:"state,omitempty"`
    StateReason *alertingv1beta1.AlertRuleStateReason `json:"stateReason,omitempty"`
    LastEvaluationTime *time.Time `json:"lastEvaluationTime,omitempty"`
    // duration of the last evaluation in seconds
    EvaluationDuration *float64 `json:"evaluationDuration,omitempty"`
    LastError *string `json:"lastError,omitempty"`
    Totals *alertingv1beta1.AlertRuleInstanceTotals `json:"totals,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AlertRuleStatus` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (alertRuleStatus *AlertRuleStatus) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AlertRuleStatus` objects.

```go
func (alertRuleStatus *AlertRuleStatus) Equals(other AlertRuleStatus) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AlertRuleStatus` fields for violations and returns them.

```go
func (alertRuleStatus *AlertRuleStatus) Validate() error
```

## See also

 * <span class="badge builder"></span> [AlertRuleStatusBuilder](./builder-AlertRuleStatusBuilder.md)
