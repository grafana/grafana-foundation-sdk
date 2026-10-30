---
title: <span class="badge builder"></span> AlertRuleStatusBuilder
---
# <span class="badge builder"></span> AlertRuleStatusBuilder

## Constructor

```go
func NewAlertRuleStatusBuilder() *AlertRuleStatusBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AlertRuleStatusBuilder) Build() (AlertRuleStatus, error)
```

### <span class="badge object-method"></span> EvaluationDuration

duration of the last evaluation in seconds

```go
func (builder *AlertRuleStatusBuilder) EvaluationDuration(evaluationDuration float64) *AlertRuleStatusBuilder
```

### <span class="badge object-method"></span> Health

```go
func (builder *AlertRuleStatusBuilder) Health(health alertingv1beta1.AlertRuleHealth) *AlertRuleStatusBuilder
```

### <span class="badge object-method"></span> LastError

```go
func (builder *AlertRuleStatusBuilder) LastError(lastError string) *AlertRuleStatusBuilder
```

### <span class="badge object-method"></span> LastEvaluationTime

```go
func (builder *AlertRuleStatusBuilder) LastEvaluationTime(lastEvaluationTime time.Time) *AlertRuleStatusBuilder
```

### <span class="badge object-method"></span> State

```go
func (builder *AlertRuleStatusBuilder) State(state alertingv1beta1.AlertRuleState) *AlertRuleStatusBuilder
```

### <span class="badge object-method"></span> StateReason

```go
func (builder *AlertRuleStatusBuilder) StateReason(stateReason alertingv1beta1.AlertRuleStateReason) *AlertRuleStatusBuilder
```

### <span class="badge object-method"></span> Totals

```go
func (builder *AlertRuleStatusBuilder) Totals(totals cog.Builder[alertingv1beta1.AlertRuleInstanceTotals]) *AlertRuleStatusBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AlertRuleStatus](./object-AlertRuleStatus.md)
