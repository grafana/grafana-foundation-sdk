---
title: <span class="badge object-type-struct"></span> RecordingRuleStatus
---
# <span class="badge object-type-struct"></span> RecordingRuleStatus

## Definition

```go
type RecordingRuleStatus struct {
    Health *alertingv1beta1.RecordingRuleHealth `json:"health,omitempty"`
    LastEvaluationTime *time.Time `json:"lastEvaluationTime,omitempty"`
    // duration of the last evaluation in seconds
    EvaluationDuration *float64 `json:"evaluationDuration,omitempty"`
    LastError *string `json:"lastError,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RecordingRuleStatus` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (recordingRuleStatus *RecordingRuleStatus) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RecordingRuleStatus` objects.

```go
func (recordingRuleStatus *RecordingRuleStatus) Equals(other RecordingRuleStatus) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RecordingRuleStatus` fields for violations and returns them.

```go
func (recordingRuleStatus *RecordingRuleStatus) Validate() error
```

## See also

 * <span class="badge builder"></span> [RecordingRuleStatusBuilder](./builder-RecordingRuleStatusBuilder.md)
