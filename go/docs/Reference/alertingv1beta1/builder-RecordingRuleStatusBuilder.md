---
title: <span class="badge builder"></span> RecordingRuleStatusBuilder
---
# <span class="badge builder"></span> RecordingRuleStatusBuilder

## Constructor

```go
func NewRecordingRuleStatusBuilder() *RecordingRuleStatusBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RecordingRuleStatusBuilder) Build() (RecordingRuleStatus, error)
```

### <span class="badge object-method"></span> EvaluationDuration

duration of the last evaluation in seconds

```go
func (builder *RecordingRuleStatusBuilder) EvaluationDuration(evaluationDuration float64) *RecordingRuleStatusBuilder
```

### <span class="badge object-method"></span> Health

```go
func (builder *RecordingRuleStatusBuilder) Health(health alertingv1beta1.RecordingRuleHealth) *RecordingRuleStatusBuilder
```

### <span class="badge object-method"></span> LastError

```go
func (builder *RecordingRuleStatusBuilder) LastError(lastError string) *RecordingRuleStatusBuilder
```

### <span class="badge object-method"></span> LastEvaluationTime

```go
func (builder *RecordingRuleStatusBuilder) LastEvaluationTime(lastEvaluationTime time.Time) *RecordingRuleStatusBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RecordingRuleStatus](./object-RecordingRuleStatus.md)
