---
title: <span class="badge object-type-struct"></span> AlertRuleSpec
---
# <span class="badge object-type-struct"></span> AlertRuleSpec

## Definition

```go
type AlertRuleSpec struct {
    Title string `json:"title"`
    Paused *bool `json:"paused,omitempty"`
    Trigger alertingv1beta1.IntervalTrigger `json:"trigger"`
    Labels map[string]alertingv1beta1.TemplateString `json:"labels,omitempty"`
    Annotations map[string]alertingv1beta1.TemplateString `json:"annotations,omitempty"`
    For *string `json:"for,omitempty"`
    KeepFiringFor *string `json:"keepFiringFor,omitempty"`
    MissingSeriesEvalsToResolve *int64 `json:"missingSeriesEvalsToResolve,omitempty"`
    NoDataState alertingv1beta1.NoDataState `json:"noDataState"`
    ExecErrState alertingv1beta1.ExecErrState `json:"execErrState"`
    NotificationSettings *alertingv1beta1.NotificationSettings `json:"notificationSettings,omitempty"`
    Expressions alertingv1beta1.ExpressionMap `json:"expressions"`
    PanelRef *alertingv1beta1.PanelRef `json:"panelRef,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AlertRuleSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (alertRuleSpec *AlertRuleSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AlertRuleSpec` objects.

```go
func (alertRuleSpec *AlertRuleSpec) Equals(other AlertRuleSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AlertRuleSpec` fields for violations and returns them.

```go
func (alertRuleSpec *AlertRuleSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [AlertRuleBuilder](./builder-AlertRuleBuilder.md)
