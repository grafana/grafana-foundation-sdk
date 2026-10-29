---
title: <span class="badge object-type-struct"></span> RecordingRuleSpec
---
# <span class="badge object-type-struct"></span> RecordingRuleSpec

## Definition

```go
type RecordingRuleSpec struct {
    Title string `json:"title"`
    Paused *bool `json:"paused,omitempty"`
    Trigger alertingv0alpha1.IntervalTrigger `json:"trigger"`
    Labels map[string]alertingv0alpha1.TemplateString `json:"labels,omitempty"`
    Metric alertingv0alpha1.MetricName `json:"metric"`
    Expressions alertingv0alpha1.ExpressionMap `json:"expressions"`
    TargetDatasourceUID alertingv0alpha1.DatasourceUID `json:"targetDatasourceUID"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RecordingRuleSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (recordingRuleSpec *RecordingRuleSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RecordingRuleSpec` objects.

```go
func (recordingRuleSpec *RecordingRuleSpec) Equals(other RecordingRuleSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RecordingRuleSpec` fields for violations and returns them.

```go
func (recordingRuleSpec *RecordingRuleSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [RecordingRuleBuilder](./builder-RecordingRuleBuilder.md)
